package app

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
	command := parseCommand([]string{"shortify", "help"})
	assert.True(t, strings.HasPrefix(command.description, "help ---"))
}

func (suite *CLISuite) TestListUsersCommand() {
	t := suite.T()
	command := parseCommand([]string{"shortify", "users", "list"})
	assert.True(t, strings.HasPrefix(command.description, "users list ---"))
}

func (suite *CLISuite) TestCreateUserCommand() {
	t := suite.T()
	command := parseCommand([]string{"shortify", "users", "create", "pseudomuto"})
	assert.True(t, strings.HasPrefix(command.description, "users create [username] ---"))
}

func (suite *CLISuite) TestResetUserPasswordCommand() {
	t := suite.T()
	command := parseCommand([]string{"shortify", "users", "resetpw", "pseudomuto"})
	assert.True(t, strings.HasPrefix(command.description, "users resetpw [username] ---"))
}
