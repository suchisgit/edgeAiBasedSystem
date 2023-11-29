package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type DockerImage struct {
	DockerHubImageName string `json:"dockerHubImageName"`
	AppName            string `json:"appName"`
	Replicas           string `json:"replicas"`
	Command            string `json:"command"`
}

type GithubPodCreation struct {
	GithubUrl        string `json:"githubUrl"`
	ExecutionCommand string `json:"executionCommand"`
}

type LocalContainer struct {
	ContainerLocation string `json:"containerLocation"`
	AppName           string `json:"appName"`
	Replicas          string `json:"noOfReps"`
	Command           string `json:"trigCmd"`
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

	//post for dockerhub image
	app.Post("/dockerImageName", func(c *fiber.Ctx) error {
		var req DockerImage
		// dockerImage := &DockerImage{}

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		dockerHubImageName := req.DockerHubImageName
		appName := req.AppName
		replicas := req.Replicas
		command := req.Command
		fmt.Print(dockerHubImageName)
		fmt.Print(appName)
		fmt.Print(replicas)
		fmt.Print(command)
		// all the pod creation logic goes here
		return c.JSON(fiber.Map{
			"success": true,
			"message": "able to pass data",
		})
	})

	//post api for local container
	app.Post("/localContainer", func(c *fiber.Ctx) error {
		var req LocalContainer

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		containerLocation := req.ContainerLocation
		appName := req.AppName
		replicas := req.Replicas
		command := req.Command
		fmt.Print(containerLocation)
		fmt.Print(appName)
		fmt.Print(replicas)
		fmt.Print(command)
		// all the pod creation logic goes here
		return c.JSON(fiber.Map{
			"success": true,
			"message": "able to pass data",
		})
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
