package main

import (
	"github.com/pseudomuto/shortify-go/shortify"
	"log"
	"net/http"
	"os"
)

const serverPort = ":8080"

func main() {
	shortify.InitializeDb()

	if !shortify.HandleCommandLine(os.Args) {
		setEncoder()

		router := shortify.NewRouter()
		log.Fatal(http.ListenAndServe(serverPort, router))
	}
}

func setEncoder() {
	if encoder := os.Getenv("SHORTIFY_ENCODER"); encoder != "" {
		shortify.SetDefaultEncoder(encoder)
	}
}
