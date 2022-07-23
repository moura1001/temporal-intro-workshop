package main

import (
	"log"
	"moura1001/temporal_intro/activities"
	"moura1001/temporal_intro/utils"
	"moura1001/temporal_intro/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	var (
		address   = utils.GoDotEnvVariable("ADDRESS")
		namespace = "default"
		taskQueue = "workshop"
	)

	// Create client
	clientOptions := client.Options{
		HostPort:  address,
		Namespace: namespace,
	}
	serviceClient, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalf("Unable to create client. Error: %v", err)
	}
	defer serviceClient.Close()

	// Create worker
	w := worker.New(serviceClient, taskQueue, worker.Options{})
	// Register workflows and activities
	w.RegisterWorkflow(workflows.Workflow)
	w.RegisterActivity(activities.Activity)

	// Start worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("Unable to start worker. Error: %v", err)
	}
}
