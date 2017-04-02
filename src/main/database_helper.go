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

	_ , err = db.Query("CREATE TABLE user (_id INTEGER," +
		" NAME VARCHAR(20) NOT NULL," +
		" status BOOLEAN NOT NULL," +
		" PRIMARY KEY (_id))")
	if err != nil {
		fmt.Println(err)
	}

	db.Query("INSERT INTO user VALUES ('Andrey',true)")
}