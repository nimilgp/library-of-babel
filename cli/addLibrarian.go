package main

import (
	"context"
	"fmt"

	"github.com/nimilgp/library-of-babel/dbLayer"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *config) addLibrarian(queries dbLayer.Queries, ctx context.Context) {
	var librarianName string
	var librarianPasswd string
	fmt.Print("Name of librarian to be added:")
	fmt.Scan(&librarianName)
	fmt.Print("Password:")
	fmt.Scan(&librarianPasswd)
	passwdhash, _ := bcrypt.GenerateFromPassword([]byte(librarianPasswd), 10)
	userParam := dbLayer.CreateNewUserParams{
		Uname:      librarianName,
		PasswdHash: string(passwdhash),
		Email:      "librarian",
		FirstName:  "librarian",
		LastName:   "librarian",
		UserType:   "librarian",
	}
	err := queries.CreateNewUser(ctx, userParam)
	if err == nil {
		fmt.Println("Added librarian succesfully")
	} else {
		fmt.Println("Failed to add librarian")
	}
}
