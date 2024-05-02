package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", "3001")); err != nil {
		log.Fatal(err)
	}
}
