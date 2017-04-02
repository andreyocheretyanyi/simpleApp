package main


import (
	_ "github.com/lib/pq"
	"database/sql"
	"os"
	"log"
	"fmt"
)


func open_db() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	fmt.Print(db)
	if err != nil {
		log.Fatal(err)
	}

	qv := "CREATE TABLE users ("+
		"id SERIAL,"+
		"age INT,"+
		"first_name VARCHAR(255),"+
		"last_name VARCHAR(255),"+
		"email TEXT"+
	");"

	_, err = db.Exec(qv)
	if err != nil {
		fmt.Println(err)
	}

}

