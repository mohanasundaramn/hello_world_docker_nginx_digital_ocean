package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/greet", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":5000")
}