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
	mux.HandleFunc("GET /sign-in/{$}", app.getSignIn)
	mux.HandleFunc("GET /sign-up/{$}", app.getSignUp)
	mux.HandleFunc("GET /profile", app.getProfile)
	mux.HandleFunc("GET /sign-out", app.getSignOut)
	mux.HandleFunc("GET /sign-out-sign-in-page", app.getSignOutSignInPage)
	mux.HandleFunc("GET /issue", app.getIssue)
	mux.HandleFunc("GET /approve", app.getApprove)
	mux.HandleFunc("GET /return", app.getReturn)
	mux.HandleFunc("GET /revoke", app.getRevoke)
	//ui updations
	mux.HandleFunc("GET /empty-section-body", app.getEmptySectionBody)
	mux.HandleFunc("GET /search-type/{type}", app.getSearchType)
	mux.HandleFunc("GET /profile-sign-out", app.getProfileSignOut)
	//data routes
	mux.HandleFunc("POST /search-book", app.postSearchBook)
	mux.HandleFunc("GET /browse-popular", app.getBrowsePopular)
	mux.HandleFunc("GET /browse-rating", app.getBrowseRating)
	mux.HandleFunc("POST /reserve/{BookID}", app.postReserveBook)
	mux.HandleFunc("POST /cancel/{ReservationID}/{Title}", app.postCancelReservation)
	mux.HandleFunc("POST /approve/{UserID}", app.postApproveUser)
	mux.HandleFunc("POST /disapprove/{UserID}", app.postDispproveUser)
	mux.HandleFunc("POST /revoke/{UserID}", app.postRevokeUser)
	mux.HandleFunc("POST /approval-list", app.postApprovalList)
	mux.HandleFunc("POST /revoke-list", app.postRevokeList)
	mux.HandleFunc("POST /members-list", app.postMembersList)
	mux.HandleFunc("POST /selected-user/{Uname}", app.postSelectedUser)
	mux.HandleFunc("POST /books-list", app.postAvailibleBooksList)
	mux.HandleFunc("POST /issue-general/{Title}/{BookID}", app.postIssueGeneralBook)
	mux.HandleFunc("POST /members-list-return", app.postMembersListReturn)
	mux.HandleFunc("POST /return-books-list/{Uname}", app.postReturnBookList)
	mux.HandleFunc("POST /return-book/{Title}/{TransactionID}", app.postReturnBook)
	mux.HandleFunc("POST /issue-reserved/{ReservationID}/{Title}", app.postIssueReservedBook)
	//action
	mux.HandleFunc("POST /sign-in", app.postSignIn)
	mux.HandleFunc("POST /sign-up", app.postSignUp)
	http.ListenAndServe(":3333", mux)
}
