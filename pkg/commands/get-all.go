package commands

import (
	"encoding/json"

	"mysql-controller/pkg/types"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	db, err := DBConnect()
	if err != nil {
		return c.SendString(err.Error())
	}

	result, err := db.Query("Select * from tasks;")
	if err != nil {
		return c.SendString(err.Error())
	}

	defer result.Close()
	var tasks []types.Task
	for result.Next() {
		var task types.Task

		err := result.Scan(&task.ID, &task.TaskName, &task.Status)
		if err != nil {
			return c.SendString(err.Error())
		}

		tasks = append(tasks, task)
	}

	responce, err := json.Marshal(tasks)
	if err != nil {
		return c.SendString(err.Error())
	}

	defer db.Close()
	return c.Send(responce)
}
