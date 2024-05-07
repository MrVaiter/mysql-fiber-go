package commands_test

import (
	"io"
	"mysql-controller/pkg/commands"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	app := fiber.New()
	assert := assert.New(t)

	app.Post("/tasks/:taskName", commands.Insert)

	req := httptest.NewRequest("POST", "/tasks/testTask", nil)
	resp, err := app.Test(req)

	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	assert.Nil(err)
	body := string(bodyBytes)

	expectedBody := "Task \"testTask\" successfully added!\n"
	assert.Equal(expectedBody, body)
}
