package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// docker run --name pgtest -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

func main() {
	const (
		createDBQuery = `
		CREATE DATABASE test
		`
		deleteDB    = `DROP DATABASE test`
		createTable = `CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			age INT,
			first_name TEXT,
			last_name TEXT
		  );`
		deleteTable = `drop schema public cascade;
						create schema public`
		insertQuery = `INSERT INTO users (age, first_name, last_name)
						VALUES (30, 'Gena', 'Tyurin');`
		selectQuery = `SELECT first_name, last_name FROM users`
	)

	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(deleteDB)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec(createDBQuery)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec(insertQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	var firstName, lastName string

	err = db.QueryRow(selectQuery).Scan(&firstName, &lastName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("recieved data: ", firstName, lastName)
	_, err = db.Exec(deleteTable)
	if err != nil {
		fmt.Println(err)
		return
	}

	//drop schema public cascade; create schema public

	/*
	    var sl []string
	   	rows, _ := db.Query("")
	   	for rows.Next(){
	   		var app string
	   		rows.Scan(&app)

	   		sl = append(sl, app)
	   	}
	*/

	/*var name string
	row := db.QueryRow(selectTable)

	err = row.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
	/*age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	…*/
}
