package handlers

import (
	"net/http"
	"practice/utils"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "signup",
	}
	utils.RenderTemplate(w, "sign.html", data)
}
