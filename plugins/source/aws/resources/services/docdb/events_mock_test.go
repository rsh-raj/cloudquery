package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildEventsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var e docdb.DescribeEventsOutput
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	e.Marker = nil
	m.EXPECT().DescribeEvents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&e,
		nil,
	)

	return services
}

func TestEvents(t *testing.T) {
	client.AwsMockTestHelper(t, Events(), buildEventsMock, client.TestOptions{})
}
