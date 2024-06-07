package main

import (
		"fmt"
		"log"
		"database/sql"

		"github.com/nimilgp/library-of-babel/pkg/models"

		_ "github.com/mattn/go-sqlite3"
)

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

		var users []models.User

		for rows.Next() {
				var u models.User
				rows.Scan(&u.User_id, &u.Uname, &u.Email)
				users = append(users, u)
		}
		for _, u := range users {
				fmt.Println(u)
		}
}
