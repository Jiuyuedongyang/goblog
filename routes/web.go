package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

func RegisterWebRoutes(r *mux.Router) {
	PC := new(controllers.PagesController)
	r.HandleFunc("/", PC.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", PC.About).Methods("GET").Name("about")
	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(PC.NotFound)

}