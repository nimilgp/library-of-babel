package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/nimilgp/library-of-babel/dbLayer"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	ctx     context.Context
	queries *dbLayer.Queries
}

func main() {
	var app application

	//db setup
	db, err := sql.Open("sqlite3", "babel.db") /////////////////////dataBase path name///////////////
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx := context.Background()
	queries := dbLayer.New(db)
	app.ctx = ctx
	app.queries = queries

	app.queries.CreateTableBooks(app.ctx)

	mux := http.NewServeMux()

	//static fileserver
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//sections routes
	mux.HandleFunc("GET /{$}", app.getRoot)
	mux.HandleFunc("GET /home/{$}", app.getHome)
	mux.HandleFunc("GET /browse/{$}", app.getBrowse)
	mux.HandleFunc("GET /search/{$}", app.getSearch)
	mux.HandleFunc("GET /notice/{$}", app.getNotice)
	mux.HandleFunc("GET /profile/{$}", app.getProfile)
	//data routes
	mux.HandleFunc("POST /search-book", app.postSearchBook)

	http.ListenAndServe(":3333", mux)
}
