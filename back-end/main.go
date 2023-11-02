package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Hello world")

	app := fiber.New()

	//health check endpoint
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":4000"))
}
