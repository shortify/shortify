package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ModelSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}

func (self *ModelSuite) TestIsNewReturnsTrueForNewRecords() {
	t := self.T()
	assert.True(t, new(model).isNew())
}

func (self *ModelSuite) TestIsNewReturnsFalseForExistingRecords() {
	t := self.T()
	model := model{Id: 1}
	assert.False(t, model.isNew())
}
