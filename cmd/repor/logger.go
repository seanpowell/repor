package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

// Logger creates server entries when accessing endpoints.
func Logger(h httprouter.Handle, name string) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		h(w, r, p)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
