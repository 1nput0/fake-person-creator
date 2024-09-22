package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1nput0/fake-person-creator/faker"
	"github.com/gorilla/mux"
)

func TestGetMaleRandomPerson(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/male-person", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getMaleRandomPerson)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var person faker.Person
	err = json.Unmarshal(rr.Body.Bytes(), &person)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if person.FirstName == "" || person.LastName == "" || person.Email == "" {
		t.Errorf("Invalid person data: %+v", person)
	}
}

func TestGetFemaleRandomPerson(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/female-person", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getFemaleRandomPerson)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var person faker.Person
	err = json.Unmarshal(rr.Body.Bytes(), &person)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if person.FirstName == "" || person.LastName == "" || person.Email == "" {
		t.Errorf("Invalid person data: %+v", person)
	}
}

func TestAPIEndpoints(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/male-person", getMaleRandomPerson).Methods("GET")
	r.HandleFunc("/api/female-person", getFemaleRandomPerson).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/male-person")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	resp, err = http.Get(ts.URL + "/api/female-person")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}
}
