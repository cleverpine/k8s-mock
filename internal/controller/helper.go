package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

func sendFile(c *fiber.Ctx, filePath string) error {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).SendString(string(fileContents))
}
