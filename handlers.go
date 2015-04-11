package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

const TokenRouteParam = "token"

func RedirectShow(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	token := params[TokenRouteParam]
	redir, err := FindRedirectByToken(token)

	if err != nil {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response.WriteHeader(http.StatusNotFound)
	} else {
		response.Header().Set("Content-Type", "text/html; charset=UTF-8")
		http.Redirect(response, request, redir.Url, http.StatusMovedPermanently)
	}
}
