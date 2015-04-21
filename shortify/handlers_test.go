package shortify

import (
	"bytes"
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
}

func (suite *HandlersSuite) SetupTest() {
	suite.redirect = NewRedirect("http://www.google.com/")
	suite.redirect.Save()
}

func (suite *HandlersSuite) TearDownSuite() {
	SetCurrentDb(false)
}

func (suite *HandlersSuite) TearDownTest() {
	TruncateDb()
}

func (suite *HandlersSuite) TestPerformRedirectHandlerWhenFound() {
	t := suite.T()
	request, _ := http.NewRequest("GET", "http://example.com/"+suite.redirect.Token, nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{token}", http.HandlerFunc(performRedirectHandler))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusMovedPermanently, response.Code)
	assert.Equal(t, suite.redirect.Url, response.HeaderMap.Get("Location"))
}

func (suite *HandlersSuite) TestPerformRedirectHandlerWhenNotFound() {
	t := suite.T()
	request, _ := http.NewRequest("GET", "http://example.com/notFound", nil)
	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/{token}", http.HandlerFunc(performRedirectHandler))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
}

func makeUser(name string, t *testing.T) *User {
	user := NewUser("testuser")
	err := user.Save()
	assert.Nil(t, err)

	return user
}

func (suite *HandlersSuite) TestCreateRedirectHandler() {
	t := suite.T()
	user := makeUser("testuser", t)
	params := []byte(`{"url":"http://www.google.com/"}`)
	request, _ := http.NewRequest("POST", "http://example.com/redirects", bytes.NewBuffer(params))
	request.SetBasicAuth(user.Name, user.Password)

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/redirects", http.HandlerFunc(createRedirectHandler))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)
}

func (suite *HandlersSuite) TestCreateRedirectHandlerWithBadParams() {
	t := suite.T()
	user := makeUser("testuser", t)
	params := []byte(`{"badParam": "http://www.google.com/"}`)
	request, _ := http.NewRequest("POST", "http://example.com/redirects", bytes.NewBuffer(params))
	request.SetBasicAuth(user.Name, user.Password)

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/redirects", http.HandlerFunc(createRedirectHandler))
	router.ServeHTTP(response, request)

	assert.Equal(t, HTTPUnprocessableEntity, response.Code)
}

func (suite *HandlersSuite) TestCreateRedirectHandlerWithBadPassword() {
	t := suite.T()
	user := makeUser("testuser", t)
	params := []byte(`{"url": "http://www.google.com/"}`)
	request, _ := http.NewRequest("POST", "http://example.com/redirects", bytes.NewBuffer(params))
	request.SetBasicAuth(user.Name, "Bad Password")

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/redirects", http.HandlerFunc(createRedirectHandler))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func (suite *HandlersSuite) TestCreateRedirectHandlerWithBadUser() {
	t := suite.T()
	user := makeUser("testuser", t)
	params := []byte(`{"url": "http://www.google.com/"}`)
	request, _ := http.NewRequest("POST", "http://example.com/redirects", bytes.NewBuffer(params))
	request.SetBasicAuth("whoisthis", user.Password)

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/redirects", http.HandlerFunc(createRedirectHandler))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusUnauthorized, response.Code)
}
