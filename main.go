package main

import (
	"log"
	"net/http"
	"practice/auth"
	"practice/handlers"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/loginAuth", auth.AuthLogin)
	http.HandleFunc("/signup", handlers.SignUpHandler)
	http.HandleFunc("/authSignUp", auth.AuthSignup)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/logout", handlers.LogOutHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/category", handlers.CategoryHandler)
	http.HandleFunc("/trending",handlers.TrendingHandler)
	http.HandleFunc("/entertainment", handlers.EntertainmentHandler)
	http.HandleFunc("/science&technolody", handlers.TechHandler)
	http.HandleFunc("/lifestyle", handlers.LifestyleHandlerHandler)
	log.Println("http server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
