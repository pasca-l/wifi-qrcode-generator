package main

import (
	"log"

	"github.com/pasca-l/wifi-qrcode-generator/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
