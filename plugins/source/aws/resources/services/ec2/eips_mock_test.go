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

func buildEc2Eips(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	a := types.Address{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	ip := "1.1.1.1"
	a.CarrierIp = &ip
	a.PublicIp = &ip
	a.CustomerOwnedIp = &ip
	a.PrivateIpAddress = &ip
	pool := "1.1.1.1/0"
	a.CustomerOwnedIpv4Pool = &pool
	a.PublicIpv4Pool = &pool

	m.EXPECT().DescribeAddresses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeAddressesOutput{
			Addresses: []types.Address{a},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2Eips(t *testing.T) {
	client.AwsMockTestHelper(t, Eips(), buildEc2Eips, client.TestOptions{})
}
