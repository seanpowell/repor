package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SessionCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	LogInUser(w, r, p)
	CreateSession(w, r, p.ByName("username"))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SessionRead(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("SessionRead")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SessionUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("SessionUpdate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SessionDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("SessionDelete")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserCreate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserRead(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := LogInUser(w, r, p)
	fmt.Println(username)
	if username != "" {
		CreateSession(w, r, username)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UserDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
