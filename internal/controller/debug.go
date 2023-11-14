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

	fmt.Printf(
		"Request URL: %s\nRequest Method: %s\nRequest Body: %s\nResponse: %d\n",
		c.OriginalURL(), c.Method(), body, c.Response().StatusCode(),
	)

	fmt.Println()
	fmt.Println()

	return err
}
