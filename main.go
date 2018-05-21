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
	todayslessons = append(todayslessons, Todayslesson{ID: "1", Name: "I am", Description: "a good student", Unit_id: "1"})
	todayslessons = append(todayslessons, Todayslesson{ID: "2", Name: "We are", Description: "old friends", Unit_id: "1"})
	todayslessons = append(todayslessons, Todayslesson{ID: "3", Name: "The pencil", Description: "is on the desk", Unit_id: "1"})

	todayslessons = append(todayslessons, Todayslesson{ID: "4", Name: "She likes", Description: "to sit  in the sun", Unit_id: "2"})
	todayslessons = append(todayslessons, Todayslesson{ID: "5", Name: "He watches", Description: "his son in the park", Unit_id: "2"})
	todayslessons = append(todayslessons, Todayslesson{ID: "6", Name: "They always play", Description: "tennis on Sundays", Unit_id: "2"})

	todayslessons = append(todayslessons, Todayslesson{ID: "7", Name: "We are", Description: "singing very nicely", Unit_id: "3"})
	todayslessons = append(todayslessons, Todayslesson{ID: "8", Name: "Are you", Description: "Metal?", Unit_id: "3"})
	todayslessons = append(todayslessons, Todayslesson{ID: "9", Name: "Bella is", Description: "playing the guitar and Andrew is listening to her", Unit_id: "3"})

	todayslessons = append(todayslessons, Todayslesson{ID: "10", Name: "You ", Description: "worked very hard last week", Unit_id: "4"})
	todayslessons = append(todayslessons, Todayslesson{ID: "11", Name: "He always", Description: "walked to school", Unit_id: "4"})
	todayslessons = append(todayslessons, Todayslesson{ID: "12", Name: "did you", Description: "paint my house last weekend", Unit_id: "4"})

	todayslessons = append(todayslessons, Todayslesson{ID: "13", Name: "She was", Description: "knocking the door", Unit_id: "5"})
	todayslessons = append(todayslessons, Todayslesson{ID: "14", Name: "They were", Description: "not playing in the park", Unit_id: "5"})
	todayslessons = append(todayslessons, Todayslesson{ID: "15", Name: "The plane was", Description: "leaving the airport", Unit_id: "5"})

	todayslessons = append(todayslessons, Todayslesson{ID: "16", Name: "I´m", Description: "bored", Unit_id: "6"})
	todayslessons = append(todayslessons, Todayslesson{ID: "17", Name: "He gave me", Description: "a broken glass.", Unit_id: "6"})
	todayslessons = append(todayslessons, Todayslesson{ID: "18", Name: "Two men were", Description: "arrested yesterday", Unit_id: "6"})

	todayslessons = append(todayslessons, Todayslesson{ID: "19", Name: "She will", Description: "read the newspaper tomorrow", Unit_id: "7"})
	todayslessons = append(todayslessons, Todayslesson{ID: "20", Name: "Will he", Description: "watch his son in the race", Unit_id: "7"})
	todayslessons = append(todayslessons, Todayslesson{ID: "21", Name: "I will", Description: "drive to work", Unit_id: "7"})

	todayslessons = append(todayslessons, Todayslesson{ID: "22", Name: "Where did you get", Description: "those boots?", Unit_id: "8"})
	todayslessons = append(todayslessons, Todayslesson{ID: "23", Name: "We can get", Description: "a beer at the bar next door", Unit_id: "8"})
	todayslessons = append(todayslessons, Todayslesson{ID: "24", Name: "When you get", Description: "to the corner, turn left", Unit_id: "8"})

	todayslessons = append(todayslessons, Todayslesson{ID: "25", Name: "If Bill studies", Description: "he will pass the exam", Unit_id: "9"})
	todayslessons = append(todayslessons, Todayslesson{ID: "26", Name: "If it doesn’t rain", Description: "we may go to the beach", Unit_id: "9"})
	todayslessons = append(todayslessons, Todayslesson{ID: "27", Name: "If Rachel had more time", Description: "she would learn to play the guitar", Unit_id: "9"})

	todayslessons = append(todayslessons, Todayslesson{ID: "28", Name: "She has", Description: "to take the children with her to Houston", Unit_id: "10"})
	todayslessons = append(todayslessons, Todayslesson{ID: "29", Name: "You have", Description: "to call her tomorrow", Unit_id: "10"})
	todayslessons = append(todayslessons, Todayslesson{ID: "30", Name: "We have", Description: "to be at the meeting next week", Unit_id: "10"})



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
