package commands

import (
	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
	db, err := DBConnect()
	if err != nil {
		return c.SendString(err.Error())
	}

	result, err := db.Query("Insert into tasks (`id`, `task_text`, `done`) values (null, '" + c.Params("taskName") + "', false);")
	if err != nil {
		return c.SendString(err.Error())
	}

	defer result.Close()
	return c.SendString("Task \"" + c.Params("taskName") + "\" successfully added!\n")
}
