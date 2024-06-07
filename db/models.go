package main

import (
		"database/sql"
)

type user struct {
		user_id int
		uname string
		passwd_hash string
		email string
		first_name string
		last_name string
		user_type string
		actions_left int
		sqltime string
		validity string
}

type book struct {
		book_id int
		title string
		author string
		year int
		genre string
		isbn int
		rating float32
		readers int
		quantity int
		sqltime string
		validity string
}

type bookState struct {
		transaction_id int
		user_id int
		book_id int
		transaction_type string
		sqltime string
		validity string
}

type reservation struct {
		reservation_id int
		user_id int
		book_id int
		sqltime string
		validity string
}

type dataBaseUnit struct {
		path string
		db *sql.DB
}
