package main

import (
	"github.com/pseudomuto/shortify-go/shortify"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if shortify.Configure(configFilePath()) {
		if !shortify.HandleCommandLine(os.Args) {
			router := shortify.NewRouter()
			log.Fatal(http.ListenAndServe(shortify.ServerPort(), router))
		}
	}
}

func configFilePath() string {
	file, _ := filepath.Abs(os.Args[0])
	return file + ".gcfg"
}
