package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
