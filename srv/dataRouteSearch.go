package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/nimilgp/library-of-babel/dbLayer"
	"golang.org/x/crypto/bcrypt"
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

func (app *application) getSearchType(w http.ResponseWriter, r *http.Request) {
	varType := r.PathValue("type")
	ts, err := template.ParseFiles("./ui/html/section-intro/selection-type.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, varType)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getBrowsePopular(w http.ResponseWriter, r *http.Request) {
	books, err := app.queries.RetrieveBooksByPopularity(app.ctx)
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

func (app *application) getBrowseRating(w http.ResponseWriter, r *http.Request) {
	books, err := app.queries.RetrieveBooksByRating(app.ctx)
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

func (app *application) postSignIn(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	passwd := r.PostFormValue("passwd")
	// fmt.Print(uname)
	passwdhash, err := app.queries.RetrievePsswdHash(app.ctx, uname)
	if err != nil {
		log.Print(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(passwdhash), []byte(passwd)) == nil {
		fmt.Println("User and Password Matches")
	} else {
		fmt.Print("Invalid")
	}
}

func (app *application) postSignUp(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	passwd := r.PostFormValue("passwd")
	bytes, _ := bcrypt.GenerateFromPassword([]byte(passwd), 14)
	passwdhash := string(bytes)
	email := r.PostFormValue("email")
	fname := r.PostFormValue("fname")
	lname := r.PostFormValue("lname")
	user := dbLayer.CreateNewUserParams{
		Uname:      uname,
		PasswdHash: passwdhash,
		Email:      email,
		FirstName:  fname,
		LastName:   lname,
		UserType:   "approvalreq",
	}
	fmt.Println(user)
	err := app.queries.CreateNewUser(app.ctx, user)
	if err != nil {
		log.Print(err)
	}
}
