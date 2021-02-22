package models

type RequestError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
