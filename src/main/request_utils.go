package main

import (
	"net/http"
	"encoding/json"
)

var users []User

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users = get_element()
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
		add_element(user.Name,user.Done)
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