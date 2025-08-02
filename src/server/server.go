package server

import (
	"net/http"
	"time"
)

func Serve() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/qrcode", qrcodeHandler)

	server := http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadHeaderTimeout: 10 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
