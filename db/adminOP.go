package main

//import "log"

func (dbu *dataBaseUnit) addLibrarian(id int) {
		const statement string = `
				UPDATE users
				SET user_type='librarian'
				WHERE user_id=? AND validity='valid';
				`

				dbu.db.Exec(statement, id)
				/*if err != nil {
						log.Fatal(err)
				}*/
}

func (dbu *dataBaseUnit) removeLibrarian(id int) {
		const statement string = `
				UPDATE users
				SET validity='invalid'
				WHERE user_id=?;
				`

				dbu.db.Exec(statement, id)
				/*if err != nil {
						log.Fatal(err)
				}*/
}
