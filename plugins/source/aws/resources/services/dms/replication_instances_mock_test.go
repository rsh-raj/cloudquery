package dms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildDmsReplicationInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDatabasemigrationserviceClient(ctrl)
	l := types.ReplicationInstance{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.ReplicationInstancePrivateIpAddress = aws.String("1.2.3.4") //nolint
	l.ReplicationInstancePrivateIpAddresses = []string{"1.2.3.4"}
	l.ReplicationInstancePublicIpAddress = aws.String("1.2.3.4") //nolint
	l.ReplicationInstancePublicIpAddresses = []string{"1.2.3.4"}
	m.EXPECT().DescribeReplicationInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&databasemigrationservice.DescribeReplicationInstancesOutput{
			ReplicationInstances: []types.ReplicationInstance{l},
		}, nil)
	lt := types.Tag{}
	if err := faker.FakeObject(&lt); err != nil {
		t.Fatal(err)
	}
	lt.ResourceArn = l.ReplicationInstanceArn
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&databasemigrationservice.ListTagsForResourceOutput{
			TagList: []types.Tag{lt},
		}, nil)
	return client.Services{
		Databasemigrationservice: m,
	}
}

func TestDmsReplicationInstances(t *testing.T) {
	client.AwsMockTestHelper(t, ReplicationInstances(), buildDmsReplicationInstances, client.TestOptions{})
}
