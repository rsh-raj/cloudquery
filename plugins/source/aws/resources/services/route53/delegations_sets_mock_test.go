package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildRoute53DelegationSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	ds := route53Types.DelegationSet{}
	if err := faker.FakeObject(&ds); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListReusableDelegationSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListReusableDelegationSetsOutput{
			DelegationSets: []route53Types.DelegationSet{ds},
		}, nil)
	return client.Services{
		Route53: m,
	}
}
func TestRoute53DelegationSets(t *testing.T) {
	client.AwsMockTestHelper(t, DelegationSets(), buildRoute53DelegationSetsMock, client.TestOptions{})
}
