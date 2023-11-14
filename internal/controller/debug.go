package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewDebugController() *Debug {
	return &Debug{}
}

type Debug struct {
}

func (ctrl *Debug) Middleware(c *fiber.Ctx) error {
	// Read the body
	body := c.Body()

	// Proceed to the next middleware
	err := c.Next()

	fmt.Printf("Request URL: %s\n", c.OriginalURL())
	fmt.Printf("Request Method: %s\n", c.Method())
	fmt.Printf("Request Body: %s\n", body)

	fmt.Printf("Response: %d\n", c.Response().StatusCode())

	fmt.Println()
	fmt.Println()

	return err
}
