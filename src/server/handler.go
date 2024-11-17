package server

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func qrcodeHandler(w http.ResponseWriter, r *http.Request) {}
