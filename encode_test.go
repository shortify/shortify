package main

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

func (self *EncodeSuite) TestEncode() {
	t := self.T()

	assert.Equal(t, "0", ShortifyEncoder.Encode(0))
	assert.Equal(t, "1B", ShortifyEncoder.Encode(99))

	assert.Equal(t, "2", UnambiguousEncoder.Encode(0))
	assert.Equal(t, "3M", UnambiguousEncoder.Encode(99))
}
