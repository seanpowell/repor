package repor

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Create",
		"POST",
		"/session/new",
		SessionCreate,
	},
	Route{
		"Read",
		"GET",
		"/session/:sessionKey",
		SessionRead,
	},
	Route{
		"Update",
		"PUT",
		"/session/:sessionKey",
		SessionUpdate,
	},
	Route{
		"Update",
		"DELETE",
		"/session/:sessionKey",
		SessionDelete,
	},
	Route{
		"Create",
		"POST",
		"/user/:username/:password",
		UserCreate,
	},
	Route{
		"Read",
		"GET",
		"/user/:username/:password",
		UserRead,
	},
	Route{
		"Update",
		"PUT",
		"/user/:username/:password",
		UserUpdate,
	},
	Route{
		"Update",
		"DELETE",
		"/user/:username/:password",
		UserDelete,
	},
}
