package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	subscription := shield.DescribeSubscriptionOutput{}
	err := faker.FakeObject(&subscription)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSubscription(gomock.Any(), gomock.Any(), gomock.Any()).Return(&subscription, nil)

	return client.Services{
		Shield: m,
	}
}

func TestSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, Subscriptions(), buildSubscriptions, client.TestOptions{})
}
