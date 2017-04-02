package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
)

var users []User

func main() {
	item := User{"Anna", true}
	users = append(users, item)
	go http.HandleFunc("/", postHandler)
	go http.HandleFunc("/add", addHandler)
	log.Println("Listening...")
	port := ":"+os.Getenv("PORT")
	//port := ":5500"
	fmt.Print(port)
	open_db()
	http.ListenAndServe(port, nil)

}





