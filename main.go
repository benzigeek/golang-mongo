package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/benzigeek/golang-mongo/controllers"
	"github.com/benzigeek/golang-mongo/database"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	type respond struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}

	resp := respond{
		StatusCode: 200,
		Message:    "Online",
	}

	json.NewEncoder(w).Encode(resp)

}

func main() {

	db := database.Init()

	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/v1/auth/register", controllers.AuthSignup(db))

	log.Fatal(http.ListenAndServe(":3005", router))

	log.Println("Started the RESTful API...")

}
