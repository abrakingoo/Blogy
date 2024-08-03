package handlers

import (
	"net/http"
	"practice/utils"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		VerifiedUser bool
	}{
		Title: "signup",
		VerifiedUser: false,
	}
	utils.RenderTemplate(w, "sign.html", data)
}
