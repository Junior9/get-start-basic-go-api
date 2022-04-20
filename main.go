package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get people")
	json.NewEncoder(w).Encode(people)
}
func CreatePeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Create people")
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get person")
	params := mux.Vars(req)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(Person{})
}
func DeletePeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("delete people")
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}

func main() {

	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Francisco", LastName: "Moraes", Address: &Address{City: "Sao Paulo", State: "Sao Paulo"}})
	people = append(people, Person{ID: "2", FirstName: "Carlos", LastName: "Moraes", Address: &Address{City: "Rio Claro", State: "Sao Paulo"}})

	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeopleEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
