package main

import (
	"goblog/app/http/middlewares"

	_ "github.com/go-sql-driver/mysql"

	"goblog/bootstrap"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
