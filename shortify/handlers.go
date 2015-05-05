package shortify

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/url"
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

func performRedirectHandler(response http.ResponseWriter, request *http.Request) {
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

func createRedirectHandler(response http.ResponseWriter, request *http.Request) {
	withAuthorization(response, request, func() {
		params, err := getCreateParams(request.Body)
		if err != nil {
			renderError(response, HTTPUnprocessableEntity, "Invalid parameters")
			return
		}

		if !isValidURL(params.Url) {
			renderError(response, HTTPUnprocessableEntity, "Invalid url")
			return
		}

		redir, err := FindOrCreateRedirect(params.Url)
		if err != nil {
			renderError(response, http.StatusInternalServerError, err.Error())
		} else {
			asJson(response, http.StatusCreated, func() {
				json.NewEncoder(response).Encode(redir)
			})
		}
	})
}

func isValidURL(inputUrl string) bool {
	url, err := url.Parse(inputUrl)
	return (err == nil && url.Scheme != "")
}

func renderError(response http.ResponseWriter, code int, message string) {
	asJson(response, code, func() {
		json.NewEncoder(response).Encode(jsonErr{Code: code, Text: message})
	})
}

func withAuthorization(response http.ResponseWriter, request *http.Request, handler func()) {
	username, password, ok := request.BasicAuth()
	if ok && IsValidUser(username, password) {
		handler()
		return
	}

	renderError(response, http.StatusUnauthorized, "Unauthorized")
}

func asJson(response http.ResponseWriter, code int, handler func()) {
	response.Header().Set("Content-Type", jsonContentType)
	response.WriteHeader(code)
	handler()
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
