package main

//import "log"

func (dbu *dataBaseUnit) addBook(b book) {
		const statement string = `
				INSERT INTO books (
						title,
						author,
						year,
						genre,
						isbn,
						rating,
						readers,
						quantity
				)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

				dbu.db.Exec(statement, b.title, b.author, b.year, b.genre, b.isbn, b.rating, b.readers, b.quantity )
				/*if err != nil {
						log.Fatal(err)
				}*/
}

func (dbu *dataBaseUnit) removeBook(id int) {
		const statement string = `
				UPDATE books
				SET validity='invalid'
				WHERE book_id=?;
				`

				dbu.db.Exec(statement, id)
				/*if err != nil {
						log.Fatal(err)
				}*/
}

func (dbu *dataBaseUnit) approveMember(id int) {
		const statement string = `
				UPDATE users
				SET user_type='member'
				WHERE user_id=? AND user_type='approvalreq';
				`

				dbu.db.Exec(statement, id)
				/*if err != nil {
						log.Fatal(err)
				}*/
}

func (dbu *dataBaseUnit) removeMember(id int) {
		const statement string = `
				UPDATE users
				SET validity='invalid'
				WHERE user_id=? and user_type='member';`

				dbu.db.Exec(statement, id)
				/*if err != nil {
						log.Fatal(err)
				}*/
}
