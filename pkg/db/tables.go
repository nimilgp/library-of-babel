package main

type User struct {
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

type Book struct {
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

type BookState struct {
		transaction_id int
		user_id int
		book_id int
		transaction_type string
		sqltime string
		validity string
}

type Reservation struct {
		reservation_id int
		user_id int
		book_id int
		sqltime string
		validity string
}
