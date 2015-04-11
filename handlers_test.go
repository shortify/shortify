package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HandlersSuite struct {
	suite.Suite
	redirect *Redirect
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlersSuite))
}

func (suite *HandlersSuite) SetupSuite() {
	SetCurrentDb(true)
	InitializeDb()

	suite.redirect = NewRedirect("http://www.google.com/")
	suite.redirect.Save()
}

func (suite *HandlersSuite) TearDownSuite() {
	TruncateDb()
	SetCurrentDb(false)
}

func (suite *HandlersSuite) TestRedirectShowWhenFound() {
	t := suite.T()
	request, _ := http.NewRequest("GET", "http://example.com/"+suite.redirect.Token, nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{token}", http.HandlerFunc(RedirectShow))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusMovedPermanently, response.Code)
	assert.Equal(t, suite.redirect.Url, response.HeaderMap.Get("Location"))
}

func (suite *HandlersSuite) TestRedirectShowWhenNotFound() {
	t := suite.T()
	request, _ := http.NewRequest("GET", "http://example.com/notFound", nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{token}", http.HandlerFunc(RedirectShow))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
}
