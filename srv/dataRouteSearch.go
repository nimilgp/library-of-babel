package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/nimilgp/library-of-babel/dbLayer"
)

func (app *application) postSearchBook(w http.ResponseWriter, r *http.Request) {
	searchStr := r.PostFormValue("search-str")
	searchOpt := r.PostFormValue("search-option")
	var books []dbLayer.Book
	var err error
	switch searchOpt {
	case "title":
		books, err = app.queries.RetrieveBooksByTitleValue(app.ctx, sql.NullString{String: searchStr, Valid: true})
		if err != nil {
			log.Print(err)
		}
	case "author":
		books, err = app.queries.RetrieveBooksByAuthorValue(app.ctx, sql.NullString{String: searchStr, Valid: true})
		if err != nil {
			log.Print(err)
		}
	case "isbn":
		books, err = app.queries.RetrieveBooksByISBN(app.ctx, searchStr)
		if err != nil {
			log.Print(err)
		}
	}
	ts, err := template.ParseFiles("./ui/html/section-body/search-result-list.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, books)
	if err != nil {
		log.Print(err)
	}
}
