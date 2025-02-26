package inspector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildInspectorFindings(t *testing.T, ctrl *gomock.Controller) client.Services {
	inspectorClient := mocks.NewMockInspectorClient(ctrl)

	finding := types.Finding{}
	err := faker.FakeObject(&finding)
	if err != nil {
		t.Fatal(err)
	}

	inspectorClient.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&inspector.ListFindingsOutput{FindingArns: []string{aws.ToString(finding.Arn)}},
		nil,
	)
	inspectorClient.EXPECT().DescribeFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&inspector.DescribeFindingsOutput{Findings: []types.Finding{finding}},
		nil,
	)

	return client.Services{
		Inspector: inspectorClient,
	}
}

func TestInspectorFindings(t *testing.T) {
	client.AwsMockTestHelper(t, Findings(), buildInspectorFindings, client.TestOptions{})
}
