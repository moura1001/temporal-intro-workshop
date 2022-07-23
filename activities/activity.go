package activities

import (
	"context"
	"errors"
	"time"

	"go.temporal.io/sdk/activity"
)

type ActivityInput struct {
	Number      int
	PartialStep int
}

type ActivityOutput struct {
	Result int
}

func Activity(ctx context.Context, input ActivityInput) (ActivityOutput, error) {

	activity.GetLogger(ctx).Info("starting activity")
	activityInfo := activity.GetInfo(ctx)

	if activityInfo.Attempt < 1 {
		return ActivityOutput{}, errors.New("attempts under 1")
	}

	time.Sleep(5 * time.Second)

	output := ActivityOutput{
		Result: input.Number * input.PartialStep,
	}

	return output, nil
}
