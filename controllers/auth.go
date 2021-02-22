package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/benzigeek/golang-mongo/models"
	"github.com/benzigeek/golang-mongo/utils"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkEmailLength(e string) bool {

	length := len(e)

	if length > 0 {
		return true
	} else {
		return false
	}

}

func checkPasswordLength(p string) bool {

	length := len(p)

	if length > 5 {
		return true
	} else {
		return false
	}

}

// AuthSingup: request handler for signup request
func AuthSignup(db *mongo.Client) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var data models.IRequestRegister

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			log.Println("Unable to decode data", err)
		}

		match, err2 := regexp.MatchString("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$", data.Data.Email)

		emailLen := checkEmailLength(data.Data.Email)

		passwdLen := checkPasswordLength(data.Data.Password)

		if err2 != nil {

			respErr := models.RequestError{
				StatusCode: 500,
				Message:    "Internal Server Error",
			}

			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(respErr)

		}

		if emailLen == true {

			if passwdLen == true {

				if match == true {

					result := models.User{}

					filter := bson.M{"email": data.Data.Email}

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

					defer cancel()

					collection := db.Database("testing").Collection("users")

					err4 := collection.FindOne(ctx, filter).Decode(&result)

					if err4 == mongo.ErrNoDocuments {

						config := utils.PasswordConfig{
							Time:    1,
							Memory:  64 * 1024,
							Threads: 4,
							KeyLen:  32,
						}

						hash, err5 := utils.GeneratePassword(&config, data.Data.Password)

						if err5 != nil {

							resp := models.RequestError{
								StatusCode: 500,
								Message:    "Internal Server Error",
							}

							w.WriteHeader(http.StatusInternalServerError)

							json.NewEncoder(w).Encode(resp)

							return

						}

						ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

						defer cancel()

						_, err6 := db.Database("testing").Collection("users").InsertOne(ctx, bson.M{"email": data.Data.Email, "password": hash})

						if err6 != nil {

							resp := models.RequestError{
								StatusCode: 500,
								Message:    "Internal Server Error",
							}

							w.WriteHeader(http.StatusInternalServerError)

							json.NewEncoder(w).Encode(resp)

							return
						}

						resp := models.RequestError{
							StatusCode: 200,
							Message:    "created Account",
						}

						json.NewEncoder(w).Encode(resp)

					} else {

						resp := models.RequestError{
							StatusCode: 400,
							Message:    "Email already used",
						}

						w.WriteHeader(http.StatusBadRequest)

						json.NewEncoder(w).Encode(resp)

					}

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
					Message:    "Invalid Password",
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

}
