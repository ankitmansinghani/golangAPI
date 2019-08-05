package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go000000 MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "admin:ankitldrpceb@tcp(database-1.ce03m0xslpgd.ap-south-1.rds.amazonaws.com:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(db.Ping())
	// defer the close till after the main function has finished
	// executing

	stmt, err := db.Prepare("CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id))")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}
	defer db.Close()

}
