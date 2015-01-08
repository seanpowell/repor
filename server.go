package repor

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
)

func main() {
	SetupEnv()
	r := NewHttpRouter()
	log.Fatal(http.ListenAndServe(":3005", context.ClearHandler(r)))
}
