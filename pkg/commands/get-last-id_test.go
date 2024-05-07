package commands_test

import (
	"mysql-controller/pkg/commands"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastID(t *testing.T) {
	assert := assert.New(t)

	lastID, err := commands.GetLastId()
	assert.Nil(err)
	assert.IsType(int64(1), lastID)
}
