package server

import "net/http"

func Serve() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/qrcode", qrcodeHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
