package main

import (
	"database/sql"
	"goblog/app/http/middlewares"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/bootstrap"
	"goblog/pkg/database"

	"net/http"
	"net/url"
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

//
//func (a Article) Delete() (rowsAffected int64, err error) {
//	rs, err := db.Exec("delete from articles where id = " + strconv.FormatInt(a.ID, 10))
//	if err != nil {
//		return 0, err
//	}
//
//	if n, _ := rs.RowsAffected(); n > 0 {
//		return n, nil
//	}
//	return 0, nil
//}
//func getRouteVariable(parameterName string, r *http.Request) string {
//	vars := mux.Vars(r)
//	return vars[parameterName]
//}

func main() {

	database.Initialize()
	db = database.DB
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
