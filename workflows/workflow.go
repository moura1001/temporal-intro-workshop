package workflows

import (
	"errors"
	"moura1001/temporal_intro/activities"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WorkflowInput struct {
	Number int
}

type WorkflowOutput struct {
	Result int
}

func Workflow(ctx workflow.Context, input WorkflowInput) (WorkflowOutput, error) {
	workflow.GetLogger(ctx).Info("starting workflow")

	if input.Number < 1 {
		return WorkflowOutput{}, errors.New("invalid number")
	}

	var (
		scheduleToStartTimeout = 5 * time.Second
		startToCloseTimeout    = 7 * time.Second
		scheduleToCloseTimeout = scheduleToStartTimeout + startToCloseTimeout
	)

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:              "workshop",
		ScheduleToCloseTimeout: scheduleToCloseTimeout,
		ScheduleToStartTimeout: scheduleToStartTimeout,
		StartToCloseTimeout:    startToCloseTimeout,
		HeartbeatTimeout:       0 * time.Second,
		WaitForCancellation:    false,
		ActivityID:             "",
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    5,
		},
	})

	result := 1
	// Define query handler
	err := workflow.SetQueryHandler(ctx, "partial_result", func() (int, error) {
		return result, nil
	})
	if err != nil {
		return WorkflowOutput{}, err
	}

	for i := 1; i <= input.Number; i++ {
		// Execute activity
		var activityOutput activities.ActivityOutput

		err = workflow.ExecuteActivity(ctx, activities.Activity,
			activities.ActivityInput{Number: result, PartialStep: i},
		).Get(ctx, &activityOutput)
		if err != nil {
			return WorkflowOutput{}, err
		}

		result = activityOutput.Result
	}

	output := WorkflowOutput{
		Result: result,
	}

	return output, nil
}
