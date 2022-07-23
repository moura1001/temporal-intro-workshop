package workflows

import (
	"moura1001/temporal_intro/activities"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

func TestWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *WorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *WorkflowTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *WorkflowTestSuite) TestSuccess() {
	s.env.RegisterWorkflow(Workflow)

	result := 1
	for i := 1; i <= 5; i++ {
		s.env.OnActivity(activities.Activity, mock.Anything,
			activities.ActivityInput{Number: result, PartialStep: i},
		).Return(activities.ActivityOutput{Result: result * i}, nil)
		result *= i
	}
	s.env.ExecuteWorkflow(Workflow, WorkflowInput{5})

	s.Require().True(s.env.IsWorkflowCompleted())
	s.Require().NoError(s.env.GetWorkflowError())

	var output WorkflowOutput
	s.Require().NoError(s.env.GetWorkflowResult(&output))

	s.Equal(120, output.Result)
}
