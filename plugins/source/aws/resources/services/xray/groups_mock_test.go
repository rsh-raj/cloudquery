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

func buildGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	test := "test"

	var group types.GroupSummary
	if err := faker.FakeObject(&group); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetGroups(
		gomock.Any(),
		&xray.GetGroupsInput{},
		gomock.Any(),
	).Return(
		&xray.GetGroupsOutput{
			Groups: []types.GroupSummary{
				group,
			},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&xray.ListTagsForResourceInput{ResourceARN: group.GroupARN},
		gomock.Any(),
	).Return(
		&xray.ListTagsForResourceOutput{
			Tags: []types.Tag{
				{
					Key:   &test,
					Value: &test,
				},
			},
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestXrayGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildGroups, client.TestOptions{})
}
