package auth

import "net/http"

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticating.."))
}