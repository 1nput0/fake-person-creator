package main

import (
	"encoding/json"
	"net/http"

	"github.com/1nput0/fake-person-creator/faker"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// defining api endpoints
	r.HandleFunc("/api/male-person", enableCORS(getMaleRandomPerson)).Methods("GET")

	// start the server
	http.ListenAndServe(":8080", r)
}

func getMaleRandomPerson(w http.ResponseWriter, r *http.Request) {
	// generate a random male person
	newMale := faker.NewMalePerson()

	// convert the person struct to json
	jsonData, err := json.Marshal(newMale)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type and write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}
