package main

import (
	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

func StartWorker(logger *zap.Logger, service workflowserviceclient.Interface) {
	// TaskListName identifies set of client workflows, activities, and workers.
	// It could be your group or client or application name.
	workerOptions := worker.Options{
		Logger:       logger,
		MetricsScope: tally.NewTestScope(TaskListName, map[string]string{}),
	}

	worker := worker.New(
		service,
		Domain,
		TaskListName,
		workerOptions,
	)

	// Let's start the worker
	err := worker.Start()
	if err != nil {
		panic("Failed to start worker")
	}

	logger.Info("Started Worker.", zap.String("worker", TaskListName))
}
