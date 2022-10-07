package signals

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

const CODE_PIN_SIGNAL = "code_pin_is_valid"

var VALID_CODE_PIN = []int{3356, 4578, 2356, 0000}

// Check if pin code is valid push event
func VerifyCreditCardCodePin(workflowID string, codePin int) error {
	// Create temporal client
	temporalClient, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("Failed to create Temporal client", err)
		return nil
	}

	// Check if code pin is valid
	codePinIsValid := false
	for _, v := range VALID_CODE_PIN {
		if codePin == v {
			codePinIsValid = true
		}
	}

	// Send signal to the client
	err = temporalClient.SignalWorkflow(context.Background(), workflowID, "", CODE_PIN_SIGNAL, codePinIsValid)
	if err != nil {
		log.Fatalln("Failed to signal client", err)
		return nil
	}

	return nil
}

// Get code pin status from queue
func ReciveCreditCartCodePin(ctx workflow.Context, signalName string) (codePinIsValid bool) {
	workflow.GetSignalChannel(ctx, signalName).Receive(ctx, &codePinIsValid)
	return
}
