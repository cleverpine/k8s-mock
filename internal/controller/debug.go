package controller

import (
	"fmt"
	"strings"

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

	var builder strings.Builder

	// Iterate over all headers
	c.Request().Header.VisitAll(func(key, value []byte) {
		// Append the header key and value to the builder
		builder.WriteString("\t- ")
		builder.WriteString(string(key))
		builder.WriteString(": ")
		builder.WriteString(string(value))
		builder.WriteString("\n") // New line for each header
	})

	// Convert the builder content to a string
	headersString := builder.String()

	fmt.Printf(
		"Request: %s %s\nRequest Headers: \n%s\nRequest Body: %s\nResponse: %d\n",
		c.Method(), c.OriginalURL(), headersString, body, c.Response().StatusCode(),
	)

	fmt.Println()
	fmt.Println()

	return err
}
