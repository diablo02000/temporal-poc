package atm

import (
	"context"
	"log"
	"temporal-poc/atm/cache"
	"temporal-poc/atm/signals"
	"time"

	"go.temporal.io/sdk/workflow"
)

func WithdrawWorkflow(ctx workflow.Context, ammountValue int) error {
	// We set our activity options with ActivtiyOptions
	// if we want to use childworkflow and if we want to set custom settings for that
	// we should use ChildWorkflowOptions like that.
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx := workflow.WithActivityOptions(ctx, ao)

	// Ask user to send code pin
	workflow.ExecuteActivity(ctx, EnterPinCode, nil).Get(ctx, nil)

	// Check if code Pin is valid
	if status := signals.ReciveCreditCartCodePin(ctx, signals.CODE_PIN_SIGNAL); !status {
		log.Println("Wrong code pin !!")
	}

	log.Printf("Code PIN is valid.")

	workflow.ExecuteActivity(ctx, GiveMoney, ammountValue).Get(ctx, nil)

	return nil
}

func ValidatePinCode(ctx workflow.Context, pinCode int) error {

}

// ////////
// Create Activities
// ///
// Ask user to enter PIN code
func EnterPinCode(ctx context.Context) error {
	log.Println("Enter PIN code.")
	return nil
}

// Check if PIN code is valid
func PinCodeIsValid(ctx context.Context, pinCode int) (bool, error) {
	// init boolean
	pinCodeIsValid := false

	// Check if pinCode is in Valide Pin code
	for _, v := range cache.ValidPinCode {
		if pinCode == v {
			pinCodeIsValid = true
		}
	}

	return pinCodeIsValid, nil
}

func GiveMoney(ctx context.Context, moneyValue int) error {
	log.Printf("withdraw %d $", moneyValue)
	return nil
}
