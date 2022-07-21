package workflows

import (
	"errors"

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

	for i := 1; i <= input.Number; i++ {
		result *= i
	}

	output := Output{
		Result: result,
	}

	return output, nil
}
