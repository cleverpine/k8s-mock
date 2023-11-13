package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New(fiber.Config{})

	// Middleware to log each request
	app.Use(func(c *fiber.Ctx) error {
		// Read the body
		body := c.Body()

		// Log the details
		fmt.Printf("URL: %s\n", c.OriginalURL())
		fmt.Printf("Method: %s\n", c.Method())
		fmt.Printf("Body: %s\n", body)

		// Proceed to the next middleware
		err := c.Next()
		fmt.Printf("Response: %d\n", c.Response().StatusCode())
		fmt.Println()
		fmt.Println()

		return err
	})

	app.Static("/", "./files")

	// Add a sample route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	err := app.ListenTLS(":7777", "ssl/certificate.crt", "ssl/private-key.pem")
	panic(err.Error())
}
