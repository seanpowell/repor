package main

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
)

func main() {
	SetupEnv()
	InitDb()
	r := NewHttpRouter()
	log.Println("Main")
	log.Fatal(http.ListenAndServe(":3001", context.ClearHandler(r)))
}
