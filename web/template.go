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
	Index   template.HTML
	Content template.HTML
}

// Render templates and write them to the output.
func render(w http.ResponseWriter, d data) {
	tpl, err := template.New("html").Parse(assets.Html["main.html"])
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(w, d)
}
