package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ev docdb.DescribeDBEngineVersionsOutput
	if err := faker.FakeObject(&ev); err != nil {
		t.Fatal(err)
	}
	ev.Marker = nil
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var parameters docdb.DescribeEngineDefaultClusterParametersOutput
	if err := faker.FakeObject(&parameters); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeEngineDefaultClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameters,
		nil,
	)

	var instanceOptions docdb.DescribeOrderableDBInstanceOptionsOutput
	if err := faker.FakeObject(&instanceOptions); err != nil {
		t.Fatal(err)
	}
	instanceOptions.Marker = nil
	m.EXPECT().DescribeOrderableDBInstanceOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&instanceOptions,
		nil,
	)

	return services
}

func TestEngineVersions(t *testing.T) {
	client.AwsMockTestHelper(t, EngineVersions(), buildEngineVersionsMock, client.TestOptions{})
}
