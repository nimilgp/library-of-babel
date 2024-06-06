package main

//import "log"

func (dbu *dataBaseUnit) applyForMembership(u user) {
		const statement string = `
				INSERT INTO books (                               
						uname,                            
						passwd_hash,                      
						email,                            
						first_name,                       
						last_name                          
				)                                                 
				VALUES (?, ?, ?, ?, ?);
				`

				dbu.db.Exec(statement, u.uname, u.passwd_hash, u.email, u.first_name, u.last_name);
				/*if err != nil {
						log.Fatal(err)
				}*/
}

