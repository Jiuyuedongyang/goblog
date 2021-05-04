package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var db *sql.DB

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

var router *mux.Router

//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
//}
//
//func aboutHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
//		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
//}
//
//func notFoundHandler(w http.ResponseWriter, r *http.Request) {
//
//	w.WriteHeader(http.StatusNotFound)
//	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
//}

type Article struct {
	Title, Body string
	ID          int64
}

func getArticleByID(id string) (Article, error) {

	article := Article{}
	query := `SELECT * FROM articles WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err

}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := getRouteVariable("id", r)

	article, err := getArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			fmt.Fprint(w, "40文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(500)
			fmt.Fprint(w, "500服务器内部错误")
		}
	} else {
		rowsAffected, err := article.Delete()
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(500)
			fmt.Fprint(w, "500服务器内部错误")
		} else {
			if rowsAffected > 0 {
				indexURL, _ := router.Get("articles.index").URL()
				http.Redirect(w, r, indexURL.String(), http.StatusFound)
			} else {
				w.WriteHeader(440)
				fmt.Fprint(w, "文章未找到")
			}
		}

	}

}

func (a Article) Delete() (rowsAffected int64, err error) {
	rs, err := db.Exec("delete from articles where id = " + strconv.FormatInt(a.ID, 10))
	if err != nil {
		return 0, err
	}

	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}
	return 0, nil
}
func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func main() {

	database.Initialize()
	db = database.DB
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")

	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
