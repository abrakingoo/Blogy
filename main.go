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
	log.Println("http server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
