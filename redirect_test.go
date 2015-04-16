package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RedirectSuite struct {
	suite.Suite
}

func TestRedirectSuite(t *testing.T) {
	suite.Run(t, new(RedirectSuite))
}

func (suite *RedirectSuite) SetupSuite() {
	SetCurrentDb(true)
	InitializeDb()
}

func (suite *RedirectSuite) TearDownSuite() {
	SetCurrentDb(false)
}

func (suite *RedirectSuite) TearDownTest() {
	TruncateDb()
}

func (suite *RedirectSuite) TestNewRedirect() {
	t := suite.T()
	redir := NewRedirect("http://www.google.com/")
	assert.Equal(t, int64(0), redir.Id)
	assert.Empty(t, redir.Token)
	assert.Equal(t, "http://www.google.com/", redir.Url)
}

func (suite *RedirectSuite) TestSaveNewRecord() {
	t := suite.T()

	redir := NewRedirect("http://www.google.com/")
	err := redir.Save()

	assert.Nil(t, err)
	assert.NotEqual(t, 0, redir.Id)
	assert.NotEmpty(t, redir.Token)
}

func (suite *RedirectSuite) TestFindByTokenWhenFound() {
	t := suite.T()

	redir := NewRedirect("https://pseudomuto.com/")
	err := redir.Save()
	assert.Nil(t, err)

	found, err := FindRedirectByToken(redir.Token)
	assert.Nil(t, err)
	assert.Equal(t, redir.Id, found.Id)
}

func (suite *RedirectSuite) TestFindByTokenWhenNotFound() {
	t := suite.T()

	_, err := FindRedirectByToken("someTokenThatDoesntExist")
	assert.NotNil(t, err)
}

func (suite *RedirectSuite) TestFindOrCreateRedirect() {
	t := suite.T()

	redir := NewRedirect("https://pseudomuto.com/")
	err := redir.Save()
	assert.Nil(t, err)

	found, err := FindOrCreateRedirect("https://pseudomuto.com/")
	assert.Nil(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, redir.Id, found.Id)
}
