package main

import (
	"github.com/pseudomuto/shortify-go/shortify"
	"log"
	"net/http"
	"os"
)

const serverPort = ":8080"

func main() {
	if !shortify.HandleCommandLine(os.Args) {
		router := shortify.NewRouter()
		log.Fatal(http.ListenAndServe(serverPort, router))
	}
}
