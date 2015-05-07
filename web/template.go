package web

import (
	"kb/assets"
)

import (
	"html/template"
	"log"
	"net/http"
)

type data struct {
	Title   string
	Content string
}

func render(w http.ResponseWriter) {
	tpl, err := template.New("html").Parse(assets.Html["main.html"])
	if err != nil {
		log.Fatalln(err)
	}

	data := data{"Knowledge Base", "Knowledges"}
	tpl.Execute(w, data)
}
