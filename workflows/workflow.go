package workflows

import (
	"errors"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Input struct {
	Number int
}

type Output struct {
	Result int
}

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting workflow")

	if input.Number < 1 {
		return Output{}, errors.New("invalid number")
	}

	result := 1
	// Define query handler
	err := workflow.SetQueryHandler(ctx, "partial_result", func() (int, error) {
		return result, nil
	})
	if err != nil {
		return Output{}, err
	}

	for i := 1; i <= input.Number; i++ {
		workflow.Sleep(ctx, 5*time.Second)

		result *= i
	}

	output := Output{
		Result: result,
	}

	return output, nil
}
