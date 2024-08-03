package auth

import (
	"net/http"
	"time"

	"practice/data"
	"practice/utils"
)

func AuthSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "error parsing form values", http.StatusInternalServerError)
		return
	}

	users, err := utils.ReadDataBase()
	if err != nil {
		if err.Error() != "user not found" {
			http.Error(w, "error fetching user data "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	newuser := data.User{
		Name:      r.FormValue("first") + " " + r.FormValue("last"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Create_at: time.Now(),
		Blogs:     []data.Blog{},
	}

	ok, err := utils.WriteToDataBase(users, newuser)
	if err != nil {
		http.Error(w, "error adding user "+err.Error(), http.StatusInternalServerError)
		return
	}

	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
