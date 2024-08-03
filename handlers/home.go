package handlers

import (
	"net/http"
	"practice/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		VerifiedUser bool
	}{
		Title: "Home",
		VerifiedUser: false,
	}
	utils.RenderTemplate(w, "index.html", data)
}
