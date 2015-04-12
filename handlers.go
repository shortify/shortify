package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

const TokenRouteParam = "token"
const HTTPUnprocessableEntity = 422

const htmlContentType = "text/html; charset=UTF-8"
const jsonContentType = "application/json; charset=UTF-8"

type createParams struct {
	Url string `json:"url"`
}

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func renderError(response http.ResponseWriter, code int, message string) {
	response.Header().Set("Content-Type", jsonContentType)
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(jsonErr{Code: code, Text: message})
}

func RedirectShow(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	token := params[TokenRouteParam]
	redir, err := FindRedirectByToken(token)

	if err != nil {
		renderError(response, http.StatusNotFound, "Not Found")
	} else {
		response.Header().Set("Content-Type", htmlContentType)
		http.Redirect(response, request, redir.Url, http.StatusMovedPermanently)
	}
}

func isAuthorized(request *http.Request) bool {
	username, password, ok := request.BasicAuth()
	return ok && IsValidUser(username, password)
}

func RedirectCreate(response http.ResponseWriter, request *http.Request) {
	if !isAuthorized(request) {
		renderError(response, http.StatusUnauthorized, "Unauthorized")
		return
	}

	createParams, err := getCreateParams(request.Body)
	if err != nil {
		renderError(response, HTTPUnprocessableEntity, "Invalid parameters")
		return
	}

	redir, err := FindOrCreateRedirect(createParams.Url)
	if err != nil {
		renderError(response, http.StatusInternalServerError, err.Error())
	} else {
		response.Header().Set("Content-Type", jsonContentType)
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(redir)
	}
}

func getCreateParams(body io.ReadCloser) (createParams, error) {
	var params createParams
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&params)

	if len(params.Url) == 0 {
		err = errors.New("Invalid parameters")
	}

	return params, err
}
