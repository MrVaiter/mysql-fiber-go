package commands

import (
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	db, err := DBConnect()
	if err != nil {
		return c.SendString(err.Error())
	}

	result, err := db.Query("Update tasks set done=" + c.Params("status") + " where id=" + c.Params("id") + ";")
	if err != nil {
		return c.SendString(err.Error())
	}

	defer result.Close()
	return c.SendString("Task #" + c.Params("id") + " successfully updated!\n")
}
