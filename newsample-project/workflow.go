package main

import (
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"time"
)

func init() {
	workflow.Register(SimpleWorkflow)
}

func SimpleWorkflow(ctx workflow.Context, value string) error {
	ao := workflow.ActivityOptions{
		// TaskList that the activity needs to be scheduled on.
		// optional: The default task list with the same name as the workflow task list.
		TaskList:               TaskListName,
		ScheduleToCloseTimeout: time.Second * 60,
		ScheduleToStartTimeout: time.Second * 60,
		StartToCloseTimeout:    time.Second * 60,
		HeartbeatTimeout:       time.Second * 10,
		WaitForCancellation:    false,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	future := workflow.ExecuteActivity(ctx, "SimpleActivity", value)
	var result string
	if err := future.Get(ctx, &result); err != nil {
		return err
	}
	workflow.GetLogger(ctx).Info("Done", zap.String("result", result))
	return nil
}
