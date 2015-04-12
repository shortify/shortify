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

func RedirectShow(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	token := params[TokenRouteParam]
	redir, err := FindRedirectByToken(token)

	if err != nil {
		response.Header().Set("Content-Type", jsonContentType)
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"})
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
	response.Header().Set("Content-Type", jsonContentType)
	encoder := json.NewEncoder(response)

	if !isAuthorized(request) {
		response.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(jsonErr{Code: http.StatusUnauthorized, Text: "Unauthorized"})
		return
	}

	createParams, err := getCreateParams(request.Body)
	if err != nil {
		response.WriteHeader(HTTPUnprocessableEntity)
		encoder.Encode(jsonErr{Code: HTTPUnprocessableEntity, Text: "Invalid parameters"})
		return
	}

	redir, err := FindOrCreateRedirect(createParams.Url)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(jsonErr{Code: http.StatusInternalServerError, Text: err.Error()})
	} else {
		response.WriteHeader(http.StatusCreated)
		encoder.Encode(redir)
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
