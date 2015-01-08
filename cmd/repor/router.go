package main

import (
	"github.com/julienschmidt/httprouter"
)

func NewHttpRouter() *httprouter.Router {
	r := httprouter.New()
	for _, route := range routes {
		var h httprouter.Handle
		h = route.Handle
		h = Logger(h, route.Name)
		r.Handle(route.Method, route.Pattern, h)
	}
	return r
}
