package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) getRoot(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/root.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/home.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getBrowse(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/browse.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSearch(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/search.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getNotice(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/notice.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSignIn(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/sign-in.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSignUp(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/sign-up.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("uname")
	uname := cookie.Value
	books, _ := app.queries.RetrieveReservedBooks(app.ctx, uname)
	ts, err := template.ParseFiles("./ui/html/main-sections/profile.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, books)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSignOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "uname",
		Value: "null",
	}
	http.SetCookie(w, &cookie)
	ts, err := template.ParseFiles("./ui/html/main-sections/sign-in.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSignOutSignInPage(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "uname",
		Value: "null",
	}
	http.SetCookie(w, &cookie)
	ts, err := template.ParseFiles("./ui/html/librarian/signin-page.html")
	if err != nil {
		log.Print(err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getProfileSignOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("<a id='profile' hx-get='/sign-in' hx-target='.section-content' hx-swap='outerHTML'>Profile</a>")))
}

func (app *application) getIssue(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/librarian/issue.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getApprove(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/librarian/approve.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getReturn(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/librarian/return.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getRevoke(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/librarian/revoke.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
