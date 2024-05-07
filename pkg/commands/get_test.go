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

func TestGet(t *testing.T) {

	assert := assert.New(t)
	app := fiber.New()

	app.Get("/tasks/:id", commands.Get)

	req := httptest.NewRequest("GET", "/tasks/1", nil)
	resp, err := app.Test(req)

	assert.Nil(err, "Should be nil")
	assert.Equal(200, resp.StatusCode, "Should be successful")

	body, err := io.ReadAll(resp.Body)
	assert.Nil(err)

	var task types.Task
	err = json.Unmarshal(body, &task)

	assert.Nil(err)
	assert.Equal(int64(1), task.ID)
}
