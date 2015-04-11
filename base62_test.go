package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Base62Suite struct {
	suite.Suite
}

func TestBase62Suite(t *testing.T) {
	suite.Run(t, new(Base62Suite))
}

func (suite *Base62Suite) TestEncode() {
	t := suite.T()
	assert.Equal(t, "0", Base62Encode(0))
	assert.Equal(t, "1B", Base62Encode(99))
}

func (suite *Base62Suite) TestDecode() {
	t := suite.T()
	assert.Equal(t, 0, Base62Decode("0"))
	assert.Equal(t, 99, Base62Decode("1B"))
}
