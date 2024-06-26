package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
	var blt bookListT
	blt.DisplayReserve = true
	cookie, err := r.Cookie("uname")
	if err != nil || cookie.Value == "null" {
		blt.DisplayReserve = false
	}
	blt.Books = books

	ts, err := template.ParseFiles("./ui/html/section-body/search-result-list.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, blt)
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

type bookListT struct {
	Books          []dbLayer.Book
	DisplayReserve bool
}

func (app *application) getBrowsePopular(w http.ResponseWriter, r *http.Request) {
	books, err := app.queries.RetrieveBooksByPopularity(app.ctx)
	if err != nil {
		log.Print(err)
	}
	var blt bookListT
	blt.DisplayReserve = true
	cookie, err := r.Cookie("uname")
	if err != nil || cookie.Value == "null" {
		blt.DisplayReserve = false
	}
	blt.Books = books

	ts, err := template.ParseFiles("./ui/html/section-body/search-result-list.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, blt)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getBrowseRating(w http.ResponseWriter, r *http.Request) {
	books, err := app.queries.RetrieveBooksByRating(app.ctx)
	if err != nil {
		log.Print(err)
	}
	var blt bookListT
	blt.DisplayReserve = true
	cookie, err := r.Cookie("uname")
	if err != nil || cookie.Value == "null" {
		blt.DisplayReserve = false
	}
	blt.Books = books
	ts, err := template.ParseFiles("./ui/html/section-body/search-result-list.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, blt)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postSignIn(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	passwd := r.PostFormValue("passwd")
	//fmt.Print(uname)
	passwdhash, err := app.queries.RetrievePsswdHash(app.ctx, uname)
	if err != nil {
		log.Print(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(passwdhash), []byte(passwd)) == nil {
		cookie := http.Cookie{
			Name:  "uname",
			Value: uname,
		}
		http.SetCookie(w, &cookie)
		_, err := app.queries.RetrieveLibrarian(app.ctx, uname)
		if err != nil {
			ts, err := template.ParseFiles("./ui/html/main-sections/profile.html")
			if err != nil {
				log.Print(err)
			}
			books, _ := app.queries.RetrieveReservedBooks(app.ctx, uname)
			err = ts.Execute(w, books)
			if err != nil {
				log.Print(err)
			}
		} else {
			ts, err := template.ParseFiles("./ui/html/librarian/librarian.html")
			if err != nil {
				log.Print(err)
			}
			err = ts.Execute(w, nil)
			if err != nil {
				log.Print(err)
			}
		}
	} else {
		ts, err := template.ParseFiles("./ui/html/main-sections/wrong-passwd.html")
		if err != nil {
			log.Print(err)
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Print(err)
		}
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
	err := app.queries.CreateNewUser(app.ctx, user)
	if err != nil {
		log.Print(err)
		ts, err := template.ParseFiles("./ui/html/main-sections/sign-up-err.html")
		if err != nil {
			log.Print(err)
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Print(err)
		}
	} else {
		const varstr string = `<main class="section-content"><section class="section-intro"><h1>Apllication Sent</h1></section></main>`
		w.Write([]byte(string(varstr)))
	}
}

func (app *application) postReserveBook(w http.ResponseWriter, r *http.Request) {
	bid := r.PathValue("BookID")
	BID, _ := strconv.ParseInt(bid, 10, 64)
	book, _ := app.queries.RetrieveBookByBID(app.ctx, BID)
	cookie, _ := r.Cookie("uname")
	uname := cookie.Value
	resv := dbLayer.CreateReservationForBookParams{
		Uname: uname,
		Title: book.Title,
	}
	err := app.queries.CreateReservationForBook(app.ctx, resv)
	if err != nil {
		fmt.Println(err)
	}
	app.queries.UpdateBookQuantityDecrease(app.ctx, BID)
	ts, err := template.ParseFiles("./ui/html/section-body/list-element.html")
	if err != nil {
		log.Print(err)
	}
	upbook, _ := app.queries.RetrieveBookByBID(app.ctx, BID)
	err = ts.Execute(w, upbook)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postCancelReservation(w http.ResponseWriter, r *http.Request) {
	rid := r.PathValue("ReservationID")
	RID, _ := strconv.ParseInt(rid, 10, 64)
	title := r.PathValue("Title")
	app.queries.UpdateReservationValidity(app.ctx, RID)
	app.queries.UpdateBookQuantityIncrease(app.ctx, title)
}

func (app *application) postApproveUser(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("UserID")
	UID, _ := strconv.ParseInt(uid, 10, 64)
	varstruc := dbLayer.UpdateUserTypeParams{
		UserType: "member",
		UserID:   UID,
	}
	err := app.queries.UpdateUserType(app.ctx, varstruc)
	if err != nil {
		fmt.Println(err)
	}
}

func (app *application) postDispproveUser(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("UserID")
	UID, _ := strconv.ParseInt(uid, 10, 64)
	err := app.queries.UpdateUserValidity(app.ctx, UID)
	if err != nil {
		fmt.Println(err)
	}
}

func (app *application) postRevokeUser(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("UserID")
	UID, _ := strconv.ParseInt(uid, 10, 64)
	err := app.queries.UpdateUserValidity(app.ctx, UID)
	if err != nil {
		fmt.Println(err)
	}
}

func (app *application) postApprovalList(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	users, _ := app.queries.RetrieveUsersThatReqApprovalLike(app.ctx, sql.NullString{String: uname, Valid: true})
	ts, err := template.ParseFiles("./ui/html/librarian/approve-list.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, users)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postRevokeList(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	users, _ := app.queries.RetrieveMembersToRevokeLike(app.ctx, sql.NullString{String: uname, Valid: true})
	ts, err := template.ParseFiles("./ui/html/librarian/revoke-list.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, users)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postMembersList(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	users, err := app.queries.RetrieveMembersToRevokeLike(app.ctx, sql.NullString{String: uname, Valid: true})
	if err != nil {
		fmt.Println(err)
	}
	ts, err := template.ParseFiles("./ui/html/librarian/members-list.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, users)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postSelectedUser(w http.ResponseWriter, r *http.Request) {
	uname := r.PathValue("Uname")
	cookie := http.Cookie{
		Name:  "selected-user",
		Value: uname,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	ts, err := template.ParseFiles("./ui/html/librarian/issue-step2.html")
	if err != nil {
		log.Print(err)
	}
	books, _ := app.queries.RetrieveReservedBooks(app.ctx, uname)
	err = ts.Execute(w, books)
	if err != nil {
		log.Print(err)
	}

}

func (app *application) postAvailibleBooksList(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	books, err := app.queries.RetrieveAvailibleBooks(app.ctx, sql.NullString{String: title, Valid: true})
	if err != nil {
		fmt.Println(err)
	}
	ts, err := template.ParseFiles("./ui/html/librarian/book-list.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, books)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postIssueGeneralBook(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("Title")
	bid := r.PathValue("BookID")
	cookie, _ := r.Cookie("selected-user")
	uname := cookie.Value
	BID, _ := strconv.ParseInt(bid, 10, 64)
	args := dbLayer.CreateTransactionParams{
		Uname:           uname,
		Title:           title,
		TransactionType: "issue",
	}
	app.queries.CreateTransaction(app.ctx, args)
	app.queries.UpdateBookQuantityDecrease(app.ctx, BID)
}

func (app *application) postMembersListReturn(w http.ResponseWriter, r *http.Request) {
	uname := r.PostFormValue("uname")
	members, _ := app.queries.RetrieveMembersToRevokeLike(app.ctx, sql.NullString{String: uname, Valid: true})
	ts, err := template.ParseFiles("./ui/html/librarian/members-list-return.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, members)
	if err != nil {
		log.Print(err)
	}

}

func (app *application) postReturnBookList(w http.ResponseWriter, r *http.Request) {
	uname := r.PathValue("Uname")
	cookie := http.Cookie{
		Name:  "selected-user",
		Value: uname,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	books, _ := app.queries.RetrieveReturnableBooksOfUser(app.ctx, uname)
	ts, err := template.ParseFiles("./ui/html/librarian/book-list-return.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, books)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) postReturnBook(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("Title")
	tid := r.PathValue("TransactionID")
	TID, _ := strconv.ParseInt(tid, 10, 64)
	cookie, _ := r.Cookie("selected-user")
	uname := cookie.Value
	app.queries.UpdateTransactionValidity(app.ctx, TID)
	app.queries.UpdateBookQuantityIncrease(app.ctx, title)
	args := dbLayer.CreateTransactionParams{
		Uname:           uname,
		Title:           title,
		TransactionType: "return",
	}
	app.queries.CreateTransaction(app.ctx, args)
}

func (app *application) postIssueReservedBook(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("Title")
	rid := r.PathValue("ReservationID")
	cookie, _ := r.Cookie("selected-user")
	uname := cookie.Value
	RID, _ := strconv.ParseInt(rid, 10, 64)
	args := dbLayer.CreateTransactionParams{
		Uname:           uname,
		Title:           title,
		TransactionType: "issue",
	}
	app.queries.CreateTransaction(app.ctx, args)
	app.queries.UpdateReservationValidity(app.ctx, RID)
}
