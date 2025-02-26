package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildRdsDbSecurityGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBSecurityGroup
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSecurityGroups(
		gomock.Any(),
		&rds.DescribeDBSecurityGroupsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSecurityGroupsOutput{DBSecurityGroups: []types.DBSecurityGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: g.DBSecurityGroupArn},
		gomock.Any(),
	).Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	return client.Services{Rds: mock}
}

func TestRDSDBSecurityGroups(t *testing.T) {
	client.AwsMockTestHelper(t, DbSecurityGroups(), buildRdsDbSecurityGroups, client.TestOptions{})
}
