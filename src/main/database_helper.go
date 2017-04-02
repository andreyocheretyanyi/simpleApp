package main


import (
	_ "github.com/lib/pq"
	"database/sql"
	"os"
	"log"
	"fmt"
)

const db_url  = "postgres://lyxswjipntxtho:c9bbfdaa77e45ca80ef725d482571153c5f73049b53b81323bf4b44f92a95087@ec2-54-228-235-185.eu-west-1.compute.amazonaws.com:5432/d9eavq5joinmbt"

func open_db() {
	db, err := sql.Open("postgres", os.Getenv(db_url))
	fmt.Print(db)
	if err != nil {
		log.Fatal(err)
	}

	_ , err = db.Exec("CREATE TABLE user (_id INTEGER," +
		" NAME VARCHAR(20) NOT NULL," +
		" status BOOLEAN NOT NULL," +
		" PRIMARY KEY (_id))")
	if err != nil {
		fmt.Println(err)
	}
}