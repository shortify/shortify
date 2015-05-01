package shortify

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EncodeSuite struct {
	suite.Suite
}

func TestEncodeSuite(t *testing.T) {
	suite.Run(t, new(EncodeSuite))
}

func (self *EncodeSuite) SetupSuite() {
	Configure("../examples/sqlite3.gcfg")
}

func (self *EncodeSuite) TestEncode() {
	t := self.T()

	assert.Equal(t, "0", shortifyEncoder.encode(0))
	assert.Equal(t, "1B", shortifyEncoder.encode(99))
}
