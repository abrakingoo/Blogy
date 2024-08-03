package handlers

import (
	"net/http"
	"practice/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Home",
	}
	utils.RenderTemplate(w, "index.html", data)
}
