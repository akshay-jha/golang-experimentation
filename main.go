package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people = []Person{Person{
	ID: "1", Firstname: "A",
	Lastname: "Z",
	Address:  &Address{City: "Bengaluru", State: "Karnataka"}},
	Person{ID: "2",
		Firstname: "B",
		Lastname:  "Y",
		Address:   &Address{City: "Hubli", State: "Karnataka"}},
	Person{ID: "3",
		Firstname: "C",
		Lastname:  "X",
		Address:   &Address{City: "Darbhanga", State: "Bihar"}}}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("Get")
	router.HandleFunc("/people/{id}", GetPerson).Methods("Get")
	http.ListenAndServe(":8000", router)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	for i := range people {
		if params["id"] == people[i].ID {
			person = people[i]
		}
	}
	json.NewEncoder(w).Encode(person)
}
