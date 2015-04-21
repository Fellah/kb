package web

import (
	"net/http"
	"time"
)

type handler struct{}

// Base handler for HTTP requests.
func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Ignore favicon.
	render(w)
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
