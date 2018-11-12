package view

import (
	"os"
	"testing"

	"github.com/researchlab/crawler/fed/model"
)

func TestSearchResultRender(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		t.Error(err)
	}
	page := model.SearchResult{}
	err = view.Render(out, page)
	if err != nil {
		t.Fatal(err)
	}
}
