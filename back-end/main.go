package main

import (
	"fmt"
	"log"

	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type DockerImage struct {
	DockerHubImage string `json:"dockerHubImage"`
	AppName        string `json:"appName"`
	Replicas       string `json:"replicas"`
	Command        string `json:"command"`
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

func createPods(name string, fileLocation string, imageTag string, replicaCount int) {
	print("Yo: ", imageTag)
	// imageTag = "sarvagya23/mnist-fix:1.0"
	servicePort := "5000" // Set your desired service port here

	// Generate unique timestamp
	timestamp := time.Now().Format("20060102150405")

	// Append timestamp for uniqueness
	deploymentName := fmt.Sprintf("%s-deployment-%s", name, timestamp)
	serviceName := fmt.Sprintf("%s-service-%s", name, timestamp)
	name = fmt.Sprintf("%s-%s", name, timestamp)
	// Pull Docker image
	dockerImage := fmt.Sprintf("%s", imageTag)
	pullCmd := exec.Command("docker", "pull", dockerImage)
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr
	fmt.Println("Pulling Docker image...")
	if err := pullCmd.Run(); err != nil {
		log.Fatalf("Failed to pull Docker image: %s", err)
	}

	// Generate deployment YAML
	deploymentYAML := []byte(fmt.Sprintf(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s
spec:
  replicas: %d
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: mnist-data
        image: %s
        command: ["python", "%s"]
`, deploymentName, replicaCount, name, name, dockerImage, fileLocation))

	deploymentFile := "deployment.yaml"
	if err := os.WriteFile(deploymentFile, deploymentYAML, 0644); err != nil {
		log.Fatalf("Failed to create deployment YAML file: %s", err)
	}
	fmt.Printf("Created %s\n", deploymentFile)

	// Generate service YAML
	serviceYAML := []byte(fmt.Sprintf(`
apiVersion: v1
kind: Service
metadata:
  name: %s
spec:
  selector:
    app: %s
  ports:
  - protocol: TCP
    port: %s
    targetPort: %s
`, serviceName, name, servicePort, servicePort))

	serviceFile := "service.yaml"
	if err := os.WriteFile(serviceFile, serviceYAML, 0644); err != nil {
		log.Fatalf("Failed to create service YAML file: %s", err)
	}
	fmt.Printf("Created %s\n", serviceFile)

	// Apply deployment and service YAML using kubectl
	fmt.Println("Applying deployment and service to Kubernetes cluster...")
	applyDeploymentCmd := exec.Command("kubectl", "apply", "-f", deploymentFile)
	applyDeploymentCmd.Stdout = os.Stdout
	applyDeploymentCmd.Stderr = os.Stderr
	if err := applyDeploymentCmd.Run(); err != nil {
		log.Fatalf("Failed to apply deployment: %s", err)
	}

	applyServiceCmd := exec.Command("kubectl", "apply", "-f", serviceFile)
	applyServiceCmd.Stdout = os.Stdout
	applyServiceCmd.Stderr = os.Stderr
	if err := applyServiceCmd.Run(); err != nil {
		log.Fatalf("Failed to apply service: %s", err)
	}

	fmt.Println("Deployment and service applied successfully!")
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

		dockerHubImage := req.DockerHubImage
		appName := req.AppName
		command := req.Command
		replicas, err := strconv.Atoi(req.Replicas)
		if err != nil {
			return err // Handle error if replicas cannot be converted to an integer
		}
		fmt.Println(dockerHubImage)
		fmt.Println(appName)
		fmt.Println(replicas)
		fmt.Println(command)

		createPods(appName, command, dockerHubImage, replicas)

		// all the pod creation logic goes here
		return c.JSON(fiber.Map{
			"success": true,
			"message": "created pods successfully",
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
