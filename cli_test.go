package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type CLISuite struct {
	suite.Suite
}

func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLISuite))
}

func (suite *CLISuite) TestHelpCommand() {
	t := suite.T()
	command := GetCLICommand([]string{"shortify", "help"})
	assert.True(t, strings.HasPrefix(command.Description, "help ---"))
}

func (suite *CLISuite) TestListUsersCommand() {
	t := suite.T()
	command := GetCLICommand([]string{"shortify", "users", "list"})
	assert.True(t, strings.HasPrefix(command.Description, "users list ---"))
}

func (suite *CLISuite) TestCreateUserCommand() {
	t := suite.T()
	command := GetCLICommand([]string{"shortify", "users", "create", "pseudomuto"})
	assert.True(t, strings.HasPrefix(command.Description, "users create [username] ---"))
}
