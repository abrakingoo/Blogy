package auth

import (
	"net/http"
	"text/template"

	"practice/data"
	"practice/utils"
	"github.com/gorilla/sessions"
)

var AuthenticatedUser *data.User
var Store = sessions.NewCookieStore([]byte("SESSION_KEY"))

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	users, err := utils.ReadDataBase()
	if err != nil {
		data := struct {
			Message  string
			Redirect string
			Inst     string
		}{
			Message:  err.Error(),
			Redirect: "/signup",
			Inst:     "Sign Up",
		}

		w.WriteHeader(http.StatusNotFound)
		tpl, tplErr := template.ParseFiles("templates/errors.html")
		if tplErr != nil {
			http.Error(w, "Error loading template: "+tplErr.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, data)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "error parsing form values", http.StatusInternalServerError)
		return
	}

	userEmail := r.FormValue("email")
	password := r.FormValue("password")

	for _, user := range users {
		if user.Email == userEmail {
			if user.Password == password {
				AuthenticatedUser = &user
				session, _ := Store.Get(r, "userSession")
				session.Values["email"] = userEmail
				err := session.Save(r, w)
				if err != nil {
					http.Error(w, "error saving session: "+err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "invalid username/password", http.StatusUnauthorized)
				return
			}
		}
	}

	tmpl, tplErr := template.ParseFiles("templates/errors.html")
	if tplErr != nil {
		http.Error(w, "Error loading template: "+tplErr.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Message  string
		Redirect string
		Inst     string
	}{
		Message:  "Account Not Found",
		Redirect: "/signup",
		Inst:     "Sign Up",
	}
	tmpl.Execute(w, data)
}
