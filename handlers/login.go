package handlers

import (
	"net/http"
	"practice/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title string
		VerifiedUser bool
	}{
		Title: "login",
		VerifiedUser: false,
	}

	utils.RenderTemplate(w, "login.html", data)
}
