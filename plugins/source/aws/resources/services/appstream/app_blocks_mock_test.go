package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildAppstreamAppBlocksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.AppBlock{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeAppBlocks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeAppBlocksOutput{
			AppBlocks: []types.AppBlock{object},
		}, nil)

	tagsOutput := appstream.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Appstream: m,
	}
}
func TestAppstreamAppBlocks(t *testing.T) {
	client.AwsMockTestHelper(t, AppBlocks(), buildAppstreamAppBlocksMock, client.TestOptions{})
}
