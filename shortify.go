package main

import (
	"github.com/shortify/shortify/app"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if app.Configure(configFilePath()) {
		if !app.HandleCommandLine(os.Args) {
			router := app.NewRouter()
			log.Fatal(http.ListenAndServe(app.ServerPort(), router))
		}
	}
}

func configFilePath() string {
	file, _ := filepath.Abs(os.Args[0])
	return file + ".gcfg"
}
