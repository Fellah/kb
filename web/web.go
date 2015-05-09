package web

import (
	"kb/assets"
	"kb/markdown"
)

import (
	"net/http"
	"strings"
	"time"
)

type handler struct{}

// Base handler for HTTP requests.
func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if strings.HasPrefix(path, "/css/") {
		path = strings.TrimLeft(path, "/css/")
		if css, ok := assets.Css[path]; ok {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
			w.Write([]byte(css))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		// TODO: Ignore favicon.
		data, _ := markdown.GetReadme()
		render(w, data)
	}
}

func Serve() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        new(handler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
