package main

import (
	"log"
	"net/http"
	"practice/handlers"
)

func main(){

	http.HandleFunc("/", handlers.HomeHandler)
	log.Println("http server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}