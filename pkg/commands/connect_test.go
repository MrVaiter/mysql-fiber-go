package commands_test

import (
	"database/sql"
	"mysql-controller/pkg/commands"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	assert := assert.New(t)

	db, err := commands.DBConnect()
	assert.Nil(err, "Should be nil")

	var mock *sql.DB
	assert.IsType(mock, db)
}
