package cloudtrail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildCloudtrailEventsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	event := types.Event{}
	err := faker.FakeObject(&event)
	if err != nil {
		t.Fatal(err)
	}

	event.CloudTrailEvent = aws.String("{}")
	m.EXPECT().LookupEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.LookupEventsOutput{
			Events: []types.Event{event},
		},
		nil,
	)

	return services
}

func TestCloudtrailEvents(t *testing.T) {
	client.AwsMockTestHelper(t, Events(), buildCloudtrailEventsMock, client.TestOptions{})
}
