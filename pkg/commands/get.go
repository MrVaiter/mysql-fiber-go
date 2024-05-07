package commands

import (
	"encoding/json"

	"mysql-controller/pkg/types"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	db, err := DBConnect()
	if err != nil {
		return c.SendString(err.Error())
	}

	result, err := db.Query("Select * from tasks where id=" + c.Params("id") + ";")
	if err != nil {
		return c.SendString(err.Error())
	}

	defer result.Close()
	result.Next()
	var task types.Task
	err = result.Scan(&task.ID, &task.TaskName, &task.Status)
	if err != nil {
		return c.SendString(err.Error())
	}

	responce, err := json.Marshal(task)
	if err != nil {
		return c.SendString(err.Error())
	}

	defer db.Close()
	return c.SendString(string(responce))
}
