package auth

import (
	"net/http"
	"text/template"

	"practice/data"
	"practice/utils"
)

var AuthenticatedUser *data.User

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
			Inst:     "SignUp",
		}

		w.WriteHeader(http.StatusNotFound)
		tpl, _ := template.ParseFiles("templates/errors.html")
		tpl.Execute(w, data)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing user data", http.StatusInternalServerError)
		return
	}

	userEmail := r.FormValue("email")
	password := r.FormValue("password")

	for _, user := range users {
		if user.Email == userEmail {
			if user.Password == password {
				AuthenticatedUser = &user
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
				return
			} else {
				http.Error(w, "invalid username/password", http.StatusUnauthorized)
				return
			}
		}
	}

	tmpl, _ := template.ParseFiles("templates/errors.html")

	data := struct {
		Message string
		Redirect string
		Inst string
	}{
		Message: "Account Not Found",
		Redirect: "/signup",
		Inst: "Sign up",
	}
	tmpl.Execute(w, data)
}
