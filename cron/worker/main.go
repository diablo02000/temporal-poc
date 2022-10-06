package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-poc/cron"
)

func main() {
	// Create Temporal client.
	// We should create one by process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "cron", worker.Options{})

	w.RegisterWorkflow(cron.PingWebsiteWorkflow)
	w.RegisterActivity(cron.PingWebsite)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
