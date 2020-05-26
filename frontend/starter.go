package main

import (
	"github.com/bairn/crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view/")) )

	http.Handle("/search", controller.CreateSearchResultHandle("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
	
}