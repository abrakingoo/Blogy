package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"practice/auth"
	"practice/data"
	"practice/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var articles data.Article

	// Retrieve the user session
	cookie, err := auth.Store.Get(r, "userSession")
	if err != nil {
		http.Error(w, "Error getting session: "+err.Error(), http.StatusInternalServerError)
		return
	}
	url := "https://newsapi.org/v2/top-headlines?country=us&apiKey=d6b4d534b08c4d4aa43339cc6943e773"

	res, err := http.Get(url)

	if err != nil {
		articles = data.Article{}
		http.Error(w, "Error Fetching Data: "+err.Error(), http.StatusInternalServerError)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		http.Error(w, "Error Reading Response Data: "+err.Error(), http.StatusInternalServerError)
	}

	// fmt.Println(string(resBody))

	err = json.Unmarshal(resBody, &articles)

	if err != nil {
		http.Error(w, "Error Unmarshalling Response Data: "+err.Error(), http.StatusInternalServerError)
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
			Trending     data.Article
		}{
			Title:        "Home",
			VerifiedUser: false, // Indicates the user is not signed in
			UserName:     "",
			Trending:     articles,
		}
		utils.RenderTemplate(w, "index.html", data)
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
		Trending     data.Article
	}{
		Title:        "Home",
		VerifiedUser: signed,
		UserName:     userName,
		Trending:     articles,
	}

	utils.RenderTemplate(w, "index.html", data)
}
