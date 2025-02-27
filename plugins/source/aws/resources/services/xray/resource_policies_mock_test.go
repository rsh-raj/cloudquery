package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildResourcePolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	var pols types.ResourcePolicy
	if err := faker.FakeObject(&pols); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListResourcePolicies(
		gomock.Any(),
		&xray.ListResourcePoliciesInput{},
		gomock.Any(),
	).Return(
		&xray.ListResourcePoliciesOutput{
			ResourcePolicies: []types.ResourcePolicy{
				pols,
			},
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestResourcePolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ResourcePolicies(), buildResourcePolicies, client.TestOptions{})
}
