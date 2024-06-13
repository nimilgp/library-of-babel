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
	db, err := sql.Open("sqlite3", "temp.db") /////////////////////dataBase path name///////////////
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
	mux.HandleFunc("GET /sign-in/{$}", app.getSignIn)
	mux.HandleFunc("GET /sign-up/{$}", app.getSignUp)
	//ui updations
	mux.HandleFunc("GET /empty-section-body", app.getEmptySectionBody)
	mux.HandleFunc("GET /search-type/{type}", app.getSearchType)
	//data routes
	mux.HandleFunc("POST /search-book", app.postSearchBook)
	mux.HandleFunc("GET /browse-popular", app.getBrowsePopular)
	mux.HandleFunc("GET /browse-rating", app.getBrowseRating)
	//action
	mux.HandleFunc("POST /sign-in", app.postSignIn)
	mux.HandleFunc("POST /sign-up", app.postSignUp)
	http.ListenAndServe(":3333", mux)
}
