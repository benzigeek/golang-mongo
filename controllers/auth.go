package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/benzigeek/golang-mongo/models"
	"github.com/julienschmidt/httprouter"
)

func AuthSignup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var data models.IRequestRegister

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Println("Unable to decode data", err)
	}

	match, err2 := regexp.MatchString("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$", data.Data.Email)

	length := len(data.Data.Email)

	if err2 != nil {

		respErr := models.RequestError{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(respErr)

	}

	if length > 0 {

		if match == true {

			type respond struct {
				Message string
			}

			resp := respond{
				Message: "works",
			}

			json.NewEncoder(w).Encode(resp)

		} else {

			respErr := models.RequestError{
				StatusCode: 400,
				Message:    "Invalid Email",
			}

			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(respErr)

		}

	} else {

		respErr := models.RequestError{
			StatusCode: 400,
			Message:    "Email is required",
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(respErr)

	}

}
