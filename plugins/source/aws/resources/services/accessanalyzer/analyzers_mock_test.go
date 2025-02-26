package accessanalyzer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildAccessAnalyzer(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAccessanalyzerClient(ctrl)
	u := types.AnalyzerSummary{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	f := types.FindingSummary{}
	if err := faker.FakeObject(&f); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAnalyzers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListAnalyzersOutput{
			Analyzers: []types.AnalyzerSummary{u},
		}, nil)

	m.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListFindingsOutput{
			Findings: []types.FindingSummary{f},
		}, nil)

	arch := types.ArchiveRuleSummary{}
	if err := faker.FakeObject(&arch); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListArchiveRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&accessanalyzer.ListArchiveRulesOutput{
			ArchiveRules: []types.ArchiveRuleSummary{arch},
		}, nil)

	return client.Services{
		Accessanalyzer: m,
	}
}

func TestAccessAnalyzerAnalyzer(t *testing.T) {
	client.AwsMockTestHelper(t, Analyzers(), buildAccessAnalyzer, client.TestOptions{})
}
