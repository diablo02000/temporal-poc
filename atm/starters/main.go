package starters

import (
	"atm/workflows"
	"context"

	"go.temporal.io/sdk/client"
)

func StartWithdrawWorkflowFunc(workflowID string, amountValue int) {
	// Create temporal client
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		panic(err)
	}
	defer c.Close()
	opt := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "worker-atm",
	}
	ctx := context.Background()
	if _, err := c.ExecuteWorkflow(ctx, opt, workflows.WithdrawWorkflow, amountValue); err != nil {
		panic(err)
	}
}
