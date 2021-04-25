package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
}
func main() {
	fmt.Fprint(1)
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}