package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildEventbridgeEventSourcesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)
	object := types.EventSource{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEventSources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListEventSourcesOutput{
			EventSources: []types.EventSource{object},
		}, nil)

	tagsOutput := eventbridge.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Eventbridge: m,
	}
}
func TestEventbridgeEventSources(t *testing.T) {
	client.AwsMockTestHelper(t, EventSources(), buildEventbridgeEventSourcesMock, client.TestOptions{})
}
