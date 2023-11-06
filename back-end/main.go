package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type DockerImage struct {
	DockerImage string `json:"dockerImage"`
}

func main() {
	fmt.Print("Hello GO back-end")

	app := fiber.New()

	// dockerImage := DockerImage{}

	//health check endpoint
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//post
	app.Post("/dockerImageName", func(c *fiber.Ctx) error {
		var req DockerImage
		// dockerImage := &DockerImage{}

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		dockerImage := req.DockerImage
		fmt.Print(dockerImage)
		return c.JSON(dockerImage)
	})
	log.Fatal(app.Listen(":4001"))
}
