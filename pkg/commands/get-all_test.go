package commands_test

import (
	"encoding/json"
	"io"
	"mysql-controller/pkg/commands"
	"mysql-controller/pkg/types"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	assert := assert.New(t)
	app := fiber.New()

	app.Get("/tasks", commands.GetAll)

	req := httptest.NewRequest("GET", "/tasks", nil)
	resp, err := app.Test(req)

	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(err)

	var tasks []types.Task
	err = json.Unmarshal(body, &tasks)
	assert.Nil(err)

	var mock []types.Task
	assert.IsType(mock, tasks)
}
