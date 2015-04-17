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

	if !processArgs() {
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

func processArgs() bool {
	args := os.Args
	if len(args) > 1 {
		command := shortify.GetCLICommand(args)
		command.Handler(args)
		return true
	}

	return false
}
