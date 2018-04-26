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

func CreateTodayslesson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var todayslesson Todayslesson
    _ = json.NewDecoder(r.Body).Decode(&todayslesson)
    todayslesson.ID = params["id"]
    todayslessons = append(todayslessons, todayslesson)
    json.NewEncoder(w).Encode(todayslessons)
}

func DeleteTodayslesson(w http.ResponseWriter, r *http.Request) {
	/*
	params := mux.Vars(r)
    for index, item := range todayslessons {
        //if item.ID == params["id"] {
    	if item.ID == id {
            todayslessons = append(todayslessons[:index], todayslessons[index+1]...)
            break
    	}
 		json.NewEncoder(w).Encode(todayslessons)
	}

	*/
	/*
	for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
    */
}

// our main function
func main() {
	router := mux.NewRouter()

	todayslessons = append(todayslessons, Todayslesson{ID: "1", Name: "Colors", Description: "Verde agache y me lo muerde"})
	todayslessons = append(todayslessons, Todayslesson{ID: "2", Name: "Clothes", Description: "T-shirt = camiseta"})
	todayslessons = append(todayslessons, Todayslesson{ID: "3", Name: "Chupelo", Description: "Papu lo chupa"})

	router.HandleFunc("/todayslessons", GetTodayslessons).Methods("GET")
	router.HandleFunc("/todayslessons/{id}", GetTodayslesson).Methods("GET")
	router.HandleFunc("/todayslessons/{id}", CreateTodayslesson).Methods("POST")
	router.HandleFunc("/todayslessons/{id}", DeleteTodayslesson).Methods("DELETE")
	fmt.Printf("Server Running")

	//log.Fatal(http.ListenAndServe(":4002", router))
	err := http.ListenAndServe(":4002", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
