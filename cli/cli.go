package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
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
	cfg.dbName = "blablabla.db" /////db name/////

	if !cfg.isCreate {
		if _, err := os.Stat(cfg.dbName); err == nil {
			fmt.Println("no creation -- operations (add,remove librarian)")
		} else {
			fmt.Println("Resources Are Missing")
			fmt.Println("Exiting...")
		}

	} else { //create
		if _, err := os.Stat(cfg.dbName); err == nil {
			cfg.resourcesExists()
		} else {
			cfg.createResources()
		}
	}
}
