package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benzigeek/golang-mongo/database"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":3005", router))

	database.Init()

}
