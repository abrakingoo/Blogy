package handlers

import (
	"net/http"
	"practice/auth"
	"practice/data"
	"practice/utils"
)

func TechHandler(w http.ResponseWriter, r *http.Request) {
	var userName string
	var blogs []data.Blog
	var verifiedUser bool

	// Check if user is authenticated
	if auth.AuthenticatedUser != nil {
		cookie, err := auth.Store.Get(r, "userSession")
		if err != nil {
			http.Error(w, "Error getting session: "+err.Error(), http.StatusInternalServerError)
			return
		}

		email, ok := cookie.Values["email"].(string)
		if ok || auth.AuthenticatedUser.Email == email {
			// Set user-specific data
			verifiedUser = true
			userName = auth.AuthenticatedUser.Name
			blogs = auth.AuthenticatedUser.Blogs
		}

	}

	// Sample data for categories
	category := Category{Name: "Science & Technology", Description: "Updates in science and technology.", Items: []string{"Innovation 1", "Research 2", "Tech 3"}}
	// Define the data to pass to the template
	data := PageData{
		Title:        "technolody",
		VerifiedUser: verifiedUser,
		Name:         userName,
		UserName:     userName,
		Blogs:        blogs,
		Category:     category,
	}

	// Render the template with the provided data
	utils.RenderTemplate(w, "index.html", data)
}
