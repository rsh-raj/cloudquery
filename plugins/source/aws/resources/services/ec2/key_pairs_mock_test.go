package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildEc2KeyPairs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.KeyPairInfo{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeKeyPairs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeKeyPairsOutput{
			KeyPairs: []types.KeyPairInfo{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2KeyPairs(t *testing.T) {
	client.AwsMockTestHelper(t, KeyPairs(), buildEc2KeyPairs, client.TestOptions{})
}
