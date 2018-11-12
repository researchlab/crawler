package main

import (
	"net/http"

	"github.com/researchlab/crawler/fed/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("fed/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler("view/template.html"))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
