package main


import (
	_ "github.com/lib/pq"
	"database/sql"
	"os"
	"log"
	"fmt"
)


func open_db() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	fmt.Print(db)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	return db
}

func create_table ()  {
	db := open_db()
	defer db.Close()
	qv := "CREATE TABLE users ("+
		"id SERIAL CONSTRAINT _id PRIMARY KEY,"+
		"name VARCHAR(255),"+
		"status BOOLEAN"+
		");"

	_, err := db.Exec(qv)
	if err != nil {
		fmt.Println(err)
	}

}

func add_element(name string,status bool)  {
	db := open_db()
	defer db.Close()
	_, err := db.Exec("INSERT INTO users (name, status)" +
		" VALUES ($1,$2);",
		name,status)
	if err != nil{
		fmt.Println(err)
	}
}

func get_element() []User  {
	var users []User
	db := open_db()
	defer db.Close()
	rows, err := db.Query("SELECT name, status" +
		" FROM users;")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var fio string
		var status bool
		err = rows.Scan(&fio,&status)
		if err != nil {
			fmt.Println(err)
		}
		usr := User{fio,status}
		users = append(users,usr)
	}
	rows.Close()

	return users
}

