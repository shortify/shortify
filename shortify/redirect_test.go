package shortify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RedirectSuite struct {
	suite.Suite
}

type TestEncoder struct {
}

func (self TestEncoder) encode(value int64) string {
	return "token"
}

func TestRedirectSuite(t *testing.T) {
	suite.Run(t, new(RedirectSuite))
}

func (suite *RedirectSuite) SetupSuite() {
	Configure("../examples/sqlite3.gcfg")
}

func (suite *RedirectSuite) TearDownTest() {
	shortifyDb.reset()
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

func (suite *RedirectSuite) TestSaveErrorsWhenTokenIsntUnique() {
	t := suite.T()

	originalEncoder := shortifyEncoder
	shortifyEncoder = TestEncoder{}

	redir := NewRedirect("http://www.google.com/")
	err := redir.Save()
	assert.Nil(t, err)

	redir = NewRedirect("http://www.google.com/")
	err = redir.Save()
	assert.NotNil(t, err)

	shortifyEncoder = originalEncoder
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
