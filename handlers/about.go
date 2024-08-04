package handlers

import (
	"net/http"
	"practice/auth"
	"practice/utils"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user session
	cookie, err := auth.Store.Get(r, "userSession")
	if err != nil {
		http.Error(w, "Error getting session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get email from session
	email, ok := cookie.Values["email"].(string)
	if !ok || email == "" {
		// Redirect to the public home page
		data := struct {
			Title        string
			VerifiedUser bool
			UserName     string
		}{
			Title:        "About",
			VerifiedUser: false, // Indicates the user is not signed in
			UserName:     "",
		}
		utils.RenderTemplate(w, "about.html", data)
		return
	}

	// If the session is valid, set user details
	var (
		signed   bool
		userName string
	)

	if auth.AuthenticatedUser != nil {
		signed = true
		userName = auth.AuthenticatedUser.Name
	}

	data := struct {
		Title        string
		VerifiedUser bool
		UserName     string
	}{
		Title:        "About",
		VerifiedUser: signed,
		UserName:     userName,
	}

	utils.RenderTemplate(w, "about.html", data)
}
