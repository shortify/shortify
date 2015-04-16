package shortify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type EncodeSuite struct {
	suite.Suite
}

func TestEncodeSuite(t *testing.T) {
	suite.Run(t, new(EncodeSuite))
}

func (self *EncodeSuite) TearDownTest() {
	SetDefaultEncoder("default")
}

func (self *EncodeSuite) TestDefaultEncoder() {
	t := self.T()
	assert.True(t, strings.HasPrefix(ShortifyEncoder.charset, "0123"))
}

func (self *EncodeSuite) TestSetDefaultEncoder() {
	t := self.T()

	SetDefaultEncoder("unambiguous")
	assert.True(t, strings.HasPrefix(ShortifyEncoder.charset, "2345"))
}

func (self *EncodeSuite) TestEncode() {
	t := self.T()

	assert.Equal(t, "0", ShortifyEncoder.Encode(0))
	assert.Equal(t, "1B", ShortifyEncoder.Encode(99))

	SetDefaultEncoder("unambiguous")
	assert.Equal(t, "2", ShortifyEncoder.Encode(0))
	assert.Equal(t, "3M", ShortifyEncoder.Encode(99))
}
