package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func (app *application) postSearchBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	booksList, err := app.queries.RetrieveBooksByTitleValue(app.ctx, sql.NullString{String: title, Valid: true})
	if err != nil {
		log.Fatal(err)
	}

	ts, err := template.ParseFiles("./ui/html/section-body/search-result-list.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, booksList)
	if err != nil {
		log.Print(err)
	}
}
