package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Todayslesson struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Unit_id			string `json:"unit_id,omitempty"`
}

var todayslessons []Todayslesson

func GetTodayslessons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todayslessons)
}

func GetTodayslesson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range todayslessons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todayslesson{})
}

// our main function
func main() {
	router := mux.NewRouter()

	//data
	todayslessons = append(todayslessons, Todayslesson{ID: "1", Name: "A blessing in disguise", Description: "a good thing that seemed bad at first", Unit_id: "1"})
	todayslessons = append(todayslessons, Todayslesson{ID: "2", Name: "A dime a dozen", Description: "Something common", Unit_id: "2"})
	todayslessons = append(todayslessons, Todayslesson{ID: "3", Name: "Beat around the bush", Description: "Avoid saying what you mean, usually because it is uncomfortable", Unit_id: "3"})

	router.HandleFunc("/todayslessons", GetTodayslessons).Methods("GET")
	router.HandleFunc("/todayslessons/", GetTodayslessons).Methods("GET")
	router.HandleFunc("/todayslessons/{id}", GetTodayslesson).Methods("GET")
	//router.HandleFunc("/todayslessons/{id}", CreateTodayslesson).Methods("POST")
	//router.HandleFunc("/todayslessons/{id}", DeleteTodayslesson).Methods("DELETE")
	fmt.Printf("Server Running")
	
	//log.Fatal(http.ListenAndServe(":4002", router))
	err := http.ListenAndServe(":4002", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
