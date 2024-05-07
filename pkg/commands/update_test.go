package commands_test

import (
	"io"
	"mysql-controller/pkg/commands"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	app := fiber.New()
	assert := assert.New(t)

	app.Post("/tasks/:id.:status", commands.Update)

	lastID, err := commands.GetLastId()
	assert.Nil(err)

	id := strconv.FormatInt(lastID, 10)

	req := httptest.NewRequest("POST", "/tasks/" + id + ".1", nil)
	resp, err := app.Test(req)

	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	assert.Nil(err)
	body := string(bodyBytes)

	expectedBody := "Task #" + id + " successfully updated!\n"
	assert.Equal(expectedBody, body)
}
