package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nimilgp/library-of-babel/dbLayer"
	"golang.org/x/crypto/bcrypt"
)

func createResources(cfg config) {
	fmt.Println("Creating Resorces...")
	//db setup
	db, err := sql.Open("sqlite3", cfg.dbName)
	if err != nil {
		log.Fatal("Failed to create DataBase")
	}
	fmt.Println("Created DataBase")
	defer db.Close()
	ctx := context.Background()
	queries := dbLayer.New(db)
	if queries.CreateTableUsers(ctx) != nil || queries.CreateTableBooks(ctx) != nil {
		log.Fatal("Failed Creation of Tables")
	}
	fmt.Println("Created Tables")

	fmt.Println("Creating ADMIN user")
	passwdhash, _ := bcrypt.GenerateFromPassword([]byte(cfg.ipPasswd), 10)
	userParam := dbLayer.CreateNewUserParams{
		Uname:      cfg.ipAdmin,
		PasswdHash: string(passwdhash),
		Email:      "admin",
		FirstName:  "admin",
		LastName:   "admin",
		UserType:   "admin",
	}
	fmt.Printf("ADMIN uname:%s\n", cfg.ipAdmin)
	fmt.Printf("ADMIN passwd:%s\n", cfg.ipPasswd)
	queries.CreateNewUser(ctx, userParam)
}
