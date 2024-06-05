package main

import (
		"database/sql"
		_ "github.com/mattn/go-sqlite3"
		"fmt"
		"log"
)

type User struct {
		uid int
		uname string
		email string
}

const file string = "lib-babel.db"
const q string = `select user_id, uname, email from users;`

func main() {
		db, err := sql.Open("sqlite3", file)
		defer db.Close()

		if err != nil {
				log.Fatal(err)	
		}
		rows, err := db.Query(q); 
		if err != nil {
				log.Fatal(err)
		}

		var users []User

		for rows.Next() {
				var u User
				rows.Scan(&u.uid, &u.uname, &u.email)
				users = append(users, u)
		}
		for _, u := range users {
				fmt.Println(u)
		}
}
