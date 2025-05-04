package server

import (
	"net/http"

	"github.com/pasca-l/wifi-qrcode-generator/qrcode"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func qrcodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST method required", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wifiSpec, err := qrcode.NewWifiSpec(r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	src := wifiSpec.Encode()

	qrCodeSpec, err := qrcode.NewQRCodeSpec(src, qrcode.L)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	qrCode, err := qrcode.NewQRCode(src, qrCodeSpec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	qrcode.DrawQRCode(w, qrCode)
}
