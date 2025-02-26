package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildIotTopicRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	lp := iot.ListTopicRulesOutput{}
	if err := faker.FakeObject(&lp); err != nil {
		t.Fatal(err)
	}
	lp.NextToken = nil
	m.EXPECT().ListTopicRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	p, err := buildRule()
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTopicRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		p, nil)

	tags := iot.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func buildRule() (*iot.GetTopicRuleOutput, error) {
	p := types.TopicRule{}
	if err := faker.FakeObject(&p); err != nil {
		return nil, err
	}
	a := types.Action{}
	if err := faker.FakeObject(&a); err != nil {
		return nil, err
	}
	p.Actions = []types.Action{
		a,
	}
	p.ErrorAction = &a
	o := iot.GetTopicRuleOutput{}
	if err := faker.FakeObject(&o); err != nil {
		return nil, err
	}
	o.Rule = &p
	return &o, nil
}

func TestIotTopicRules(t *testing.T) {
	client.AwsMockTestHelper(t, TopicRules(), buildIotTopicRules, client.TestOptions{})
}
