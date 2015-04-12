package main

import (
	"log"
	"net/http"
	"os"
)

const serverPort = ":8080"

func main() {
	InitializeDb()

	if !processArgs() {
		routes := []Route{
			Route{"PerformRedirect", "GET", "/{token}", RedirectShow},
			Route{"CreateRedirect", "POST", "/redirects", RedirectCreate},
		}

		router := NewRouter(routes)
		log.Fatal(http.ListenAndServe(serverPort, router))
	}
}

func processArgs() bool {
	args := os.Args
	if len(args) > 1 {
		command := GetCLICommand(args)
		command.Handler(args)
		return true
	}

	return false
}
