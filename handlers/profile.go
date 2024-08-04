package handlers

import (
	"net/http"

	"practice/auth"
	"practice/data"
	"practice/utils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is authenticated
	if auth.AuthenticatedUser == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	cookie, err := auth.Store.Get(r, "userSession")
	if err != nil {
		http.Error(w, "error getting session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	email, ok := cookie.Values["email"].(string)

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if auth.AuthenticatedUser.Email != email {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Prepare the data for rendering the template
	data := struct {
		Title        string
		VerifiedUser bool
		Name         string
		UserName     string
		Blogs        []data.Blog
	}{
		Title:        "Profile",
		VerifiedUser: true, // Adjust as needed based on user verification
		Name:         auth.AuthenticatedUser.Name,
		UserName:     auth.AuthenticatedUser.Name,
		Blogs:        auth.AuthenticatedUser.Blogs,
	}

	// Render the profile template
	utils.RenderTemplate(w, "profile.html", data)
}
