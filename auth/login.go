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

func init() {
	// Set cookie options
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,                 // 1 hour
		Secure:   true,                 // only send cookies over HTTPS
		HttpOnly: true,                 // prevent JavaScript access to cookies
		SameSite: http.SameSiteLaxMode, // prevent CSRF, adjust as needed
	}
}

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
		http.Error(w, "Error parsing form values", http.StatusInternalServerError)
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
					http.Error(w, "Error saving session: "+err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
				return
			} else {

				data := struct {
					Title        string
					VerifiedUser bool
					ErrorMessage string
				}{
					Title:        "login",
					VerifiedUser: false,
					ErrorMessage: "Invalid username or password",
				}

				utils.RenderTemplate(w, "login.html", data)
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
