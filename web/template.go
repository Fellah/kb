package web

import (
	"html/template"
	"net/http"
)

type data struct {
	Title   string
	Content string
}

var tpl *template.Template

var (
	tplMain = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8"/>
		<title>{{.Title}}</title>
	</head>

	<body>
		{{template "Index"}}
		{{.Content}}
	</body>
</html>
`

	tplIndex = `
{{define "Index"}}
<div>Index</div>
{{end}}
`
)

func init() {
	tpl, _ = template.New("html").Parse(tplMain)
	tpl.Parse(tplIndex)
}

func render(w http.ResponseWriter) {
	data := data{"Knowledge Base", "Knowledges"}
	tpl.Execute(w, data)
}
