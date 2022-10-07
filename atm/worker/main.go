package main

import (
	"log"
	"temporal-poc/atm"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	w := worker.New(c, "withdraw-atm", worker.Options{})

	w.RegisterWorkflow(atm.WithdrawWorkflow)
	w.RegisterActivity(atm.EnterPinCode)
	w.RegisterActivity(atm.PinCodeIsValid)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start withdraw worker", err)
	}

}
