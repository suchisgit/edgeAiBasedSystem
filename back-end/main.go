package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type DockerImage struct {
	DockerImage string `json:"dockerImage"`
}

type GithubPodCreation struct {
	GithubUrl        string `json:"githubUrl"`
	ExecutionCommand string `json:"executionCommand"`
}

func main() {
	fmt.Print("Hello GO back-end")

	app := fiber.New()

	// dockerImage := DockerImage{}

	app.Use(cors.New())

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

	//post
	app.Post("/githubPodDetails", func(c *fiber.Ctx) error {
		var req GithubPodCreation

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		url := req.GithubUrl
		cmd := req.ExecutionCommand

		fmt.Print("called githubpoddetials post api")
		fmt.Print(url)
		fmt.Print(cmd)

		// createPod function will be called here

		// If the pod was created successfully, return a success response
		return c.JSON(fiber.Map{
			"success": true,
			"message": "able to pass data",
		})
	})

	// post upload file

	log.Fatal(app.Listen(":4001"))
}
