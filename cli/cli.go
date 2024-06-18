package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nimilgp/library-of-babel/dbLayer"
	"golang.org/x/crypto/bcrypt"
)

type config struct {
	isCreate bool
	ipAdmin  string
	ipPasswd string
	dbName   string
}

func main() {
	var cfg config
	flag.BoolVar(&cfg.isCreate, "create", false, "Create all required resources")
	flag.StringVar(&cfg.ipAdmin, "a", "admin", "Name of admin")
	flag.StringVar(&cfg.ipPasswd, "p", "passwd", "Passwd of admin")
	flag.Parse()
	cfg.dbName = "babel.db" /////db name/////

	if !cfg.isCreate {
		if _, err := os.Stat(cfg.dbName); err != nil {
			fmt.Println("Resources Are Missing")
			fmt.Println("Exiting...")
		} else {
			db, err := sql.Open("sqlite3", cfg.dbName)
			if err != nil {
				log.Fatal("Failed to Open DataBase")
			}
			defer db.Close()
			ctx := context.Background()
			queries := dbLayer.New(db)
			cmppasswdhash, _ := queries.RetrievePsswdHash(ctx, cfg.ipAdmin)
			if bcrypt.CompareHashAndPassword([]byte(cmppasswdhash), []byte(cfg.ipPasswd)) != nil {
				fmt.Println("Admin and Passwd dont match")
				fmt.Println("Exiting...")
				return
			}
			var action int
			for {
				fmt.Println("(1)Add Librarian (2)Remove Librarian (3)Exit")
				fmt.Print("Choose:")
				fmt.Scan(&action)
				if action == 1 {
					cfg.addLibrarian(*queries, ctx)
				} else if action == 2 {
					var librarianName string
					fmt.Print("Name of librarian to be removed:")
					fmt.Scan(&librarianName)
					err := queries.InvalidateUser(ctx, librarianName)
					if err == nil {
						fmt.Println("Removed librarian succesfully")
					} else {
						fmt.Println("Failed to remove librarian")
					}
				} else {
					fmt.Println("Exiting...")
					return
				}
				fmt.Println()
			}
		}
	} else { //create
		if _, err := os.Stat(cfg.dbName); err == nil {
			cfg.resourcesExists()
		} else {
			cfg.createResources()
		}
	}
}
