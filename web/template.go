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
	Content template.HTML
}

func render(w http.ResponseWriter, d []byte) {
	tpl, err := template.New("html").Parse(assets.Html["main.html"])
	if err != nil {
		log.Fatalln(err)
	}

	data := data{"Knowledge Base", template.HTML(string(d))}
	tpl.Execute(w, data)
}
