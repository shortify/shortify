package main

import (
	"log"
	"net/http"
)

const serverPort = ":8080"

func main() {
	routes := []Route{
		Route{"PerformRedirect", "GET", "/{token}", RedirectShow},
		Route{"CreateRedirect", "POST", "/redirects", RedirectCreate},
	}

	router := NewRouter(routes)
	log.Fatal(http.ListenAndServe(serverPort, router))
}
