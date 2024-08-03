package handlers

import (
	"net/http"
	"practice/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	data := struct{
		Title string
	}{
		Title: "login",
	}
	utils.RenderTemplate(w, "login.html", data)
}