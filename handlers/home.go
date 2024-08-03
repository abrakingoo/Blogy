package handlers

import (
    "net/http"
    "text/template"
)

var templates *template.Template

func init() {
    templates = template.Must(template.ParseGlob("templates/*.html"))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title string
    }{
        Title: "Home",
    }
    renderTemplate(w, "index.html", data)
}
