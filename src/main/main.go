package main

import (
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
)

type ResponseGet struct {
	Status bool   `json:"status"`
	Users  []User `json:"users"`
}

type ResponsePost struct {
	Status bool `json:"status"`
}

type User struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var users []User

func main() {
	item := User{"Anna", true}
	users = append(users, item)
	http.HandleFunc("/", postHandler)
	http.HandleFunc("/add", addHandler)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}



func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		response := ResponseGet{true, users}
		productJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productJson)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		user := User{}
		err := decoder.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		users = append(users, user)
		response := ResponsePost{true}
		productJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productJson)
	} else {
		//productJson, _ := json.Marshal(users)
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		//w.Write(productJson)
	}
	defer r.Body.Close()

}

