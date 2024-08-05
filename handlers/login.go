package handlers

import (
	"net/http"
	"practice/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Title        string
		VerifiedUser bool
		ErrorMessage string
	}{
		Title:        "login",
		VerifiedUser: false,
		ErrorMessage: "",
	}

	utils.RenderTemplate(w, "login.html", data)
}
