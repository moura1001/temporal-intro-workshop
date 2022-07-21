package workflows

import "go.temporal.io/sdk/workflow"

func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting workflow")

	return nil
}
