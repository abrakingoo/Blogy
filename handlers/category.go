package handlers

import (
	"net/http"
	"practice/auth"
	"practice/data"
	"practice/utils"
)

type Category struct {
	Name        string
	Description string
	Items       []string
}

type PageData struct {
	Title        string
	VerifiedUser bool
	Name         string
	UserName     string
	Blogs        []data.Blog
	Category     Category
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
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

	// Parse the form data
	r.ParseForm()
	categoryName := r.FormValue("category")

	// Sample data for categories
	categories := map[string]Category{
		"trending":      {Name: "Trending", Description: "Current trending topics.", Items: []string{"Topic 1", "Topic 2", "Topic 3"}},
		"entertainment": {Name: "Entertainment", Description: "Latest in entertainment.", Items: []string{"Movie 1", "Show 2", "Music 3"}},
		"science":       {Name: "Science & Technology", Description: "Updates in science and technology.", Items: []string{"Innovation 1", "Research 2", "Tech 3"}},
		"lifestyle":     {Name: "Lifestyle", Description: "Tips and trends in lifestyle.", Items: []string{"Health 1", "Fitness 2", "Travel 3"}},
	}

	// Fetch the selected category data
	category, ok := categories[categoryName]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Define the data to pass to the template
	data := PageData{
		Title:        "Category Page",
		VerifiedUser: verifiedUser,
		Name:         userName,
		UserName:     userName,
		Blogs:        blogs,
		Category:     category,
	}

	// Render the template with the provided data
	utils.RenderTemplate(w, "index.html", data)
}
