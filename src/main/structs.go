package main


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
