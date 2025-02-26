package applicationautoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildApplicationAutoscalingPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApplicationautoscalingClient(ctrl)
	services := client.Services{
		Applicationautoscaling: m,
	}
	c := types.ScalingPolicy{}
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	output := &applicationautoscaling.DescribeScalingPoliciesOutput{
		ScalingPolicies: []types.ScalingPolicy{c},
	}
	m.EXPECT().DescribeScalingPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		output,
		nil,
	)

	return services
}

func TestApplicationAutoscalingPolicies(t *testing.T) {
	client.AllNamespaces = []string{"test-namespace"} // Just one

	client.AwsMockTestHelper(t, Policies(), buildApplicationAutoscalingPoliciesMock, client.TestOptions{})
}
