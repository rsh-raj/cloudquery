package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildWAFV2ManagedRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafv2Client(ctrl)
	var tempDescribeManagedRuleGroup wafv2.DescribeManagedRuleGroupOutput
	if err := faker.FakeObject(&tempDescribeManagedRuleGroup); err != nil {
		t.Fatal(err)
	}
	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempManagedRuleGroupSum := types.ManagedRuleGroupSummary{}
		if err := faker.FakeObject(&tempManagedRuleGroupSum); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListAvailableManagedRuleGroups(gomock.Any(), &wafv2.ListAvailableManagedRuleGroupsInput{
			Scope: scope,
		}, gomock.Any()).Return(&wafv2.ListAvailableManagedRuleGroupsOutput{
			ManagedRuleGroups: []types.ManagedRuleGroupSummary{tempManagedRuleGroupSum},
		}, nil)
		m.EXPECT().DescribeManagedRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tempDescribeManagedRuleGroup, nil)
	}

	return client.Services{Wafv2: m}
}

func TestWafV2ManagedRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ManagedRuleGroups(), buildWAFV2ManagedRuleGroupsMock, client.TestOptions{})
}
