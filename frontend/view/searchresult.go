package view

import (
	"github.com/bairn/crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template:template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render (w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}