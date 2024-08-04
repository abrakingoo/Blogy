package handlers

import (
	"net/http"
	"practice/auth"
	"practice/utils"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session from the request
	session, err := auth.Store.Get(r, "userSession")
	if err != nil {
		http.Error(w, "error getting session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear the session
	session.Options.MaxAge = -1
	session.Values = nil

	// Save the session
	if err := session.Save(r, w); err != nil {
		http.Error(w, "error saving session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title        string
		VerifiedUser bool
		UserName     string
	}{
		Title:        "Home",
		VerifiedUser: false,
		UserName:     "",
	}

	utils.RenderTemplate(w, "index.html", data)
}
