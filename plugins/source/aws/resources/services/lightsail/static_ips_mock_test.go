package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildStaticIps(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var ips lightsail.GetStaticIpsOutput
	if err := faker.FakeObject(&ips); err != nil {
		t.Fatal(err)
	}
	ips.NextPageToken = nil

	mock.EXPECT().GetStaticIps(gomock.Any(), &lightsail.GetStaticIpsInput{}, gomock.Any()).Return(&ips, nil)

	return client.Services{Lightsail: mock}
}

func TestStaticIps(t *testing.T) {
	client.AwsMockTestHelper(t, StaticIps(), buildStaticIps, client.TestOptions{})
}
