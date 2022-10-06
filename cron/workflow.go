package cron

import (
	"context"
	"log"
	"time"

	"github.com/go-ping/ping"
	"go.temporal.io/sdk/workflow"
)

// CronResult is used to return data from one cron run to the next
type CronResult struct {
	RunTime time.Time
}

// PingWebsiteWorkflow executes on the given schedule
// The schedule is provided when starting the Workflow
func PingWebsiteWorkflow(ctx workflow.Context, website string) (*CronResult, error) {

	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	thisRunTime := workflow.Now(ctx)

	err := workflow.ExecuteActivity(ctx1, PingWebsite, website).Get(ctx, nil)
	if err != nil {
		// Cron job failed
		// Next cron will still be scheduled by the Server
		workflow.GetLogger(ctx).Error("Cron job failed.", "Error", err)
		return nil, err
	}

	return &CronResult{RunTime: thisRunTime}, nil
}

// Ping website in Activity
func PingWebsite(ctx context.Context, website string) error {
	// In Activity we can query database, call external API, or do any other non-deterministic action.
	// activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)
	pinger, err := ping.NewPinger(website)
	if err != nil {
		log.Fatalln("Failed to create pinger", "Error", err)
	}

	// send uniq ping request
	pinger.Count = 1

	err = pinger.Run() // Blocks until finished.
	if err != nil {
		log.Fatalln("Failed to send ping request", "Error", err)
	}

	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	log.Printf("Ping to %s succeed in %v.", pinger.Addr(), stats.AvgRtt)

	return nil
}
