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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
		`user("id" SERIAL PRIMARY KEY,` +
		`"name" varchar(50), "status" boolean)`)
	if err != nil {
		fmt.Println(err)
	}

}