package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func (app *application) postSearchBook(w http.ResponseWriter, r *http.Request) {
	searchStr := r.PostFormValue("search-str")
	searchOpt := r.PostFormValue("search-option")
	switch searchOpt {
	case "title":
		app.searchCaseTitle(w, searchStr)
	case "author":
		app.searchCaseAuthor(w, searchStr)
	case "isbn":
		app.searchCaseISBN(w, searchStr)
	}

}

func (app *application) searchCaseTitle(w http.ResponseWriter, searchStr string) {
	books, err := app.queries.RetrieveBooksByTitleValue(app.ctx, sql.NullString{String: searchStr, Valid: true})
	if err != nil {
		log.Print(err)
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

func (app *application) searchCaseAuthor(w http.ResponseWriter, searchStr string) {
	books, err := app.queries.RetrieveBooksByAuthorValue(app.ctx, sql.NullString{String: searchStr, Valid: true})
	if err != nil {
		log.Print(err)
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

func (app *application) searchCaseISBN(w http.ResponseWriter, searchStr string) {
	books, err := app.queries.RetrieveBooksByISBN(app.ctx, searchStr)
	if err != nil {
		log.Print(err)
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
