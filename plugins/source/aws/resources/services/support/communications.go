package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func communications() *schema.Table {
	return &schema.Table{
		Name:        "aws_support_case_communications",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCommunications.html`,
		Resolver:    fetchCommunications,
		Transform:   transformers.TransformWithStruct(&types.Communication{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func fetchCommunications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Support
	p := parent.Item.(types.CaseDetails)
	input := support.DescribeCommunicationsInput{MaxResults: aws.Int32(100), CaseId: p.CaseId}

	paginator := support.NewDescribeCommunicationsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Communications
	}

	return nil
}

func mockCommunications(parent types.CaseDetails, m *mocks.MockSupportClient) error {
	communications := []types.Communication{}
	err := faker.FakeObject(&communications)
	if err != nil {
		return err
	}

	input := support.DescribeCommunicationsInput{MaxResults: aws.Int32(100), CaseId: parent.CaseId}
	m.EXPECT().DescribeCommunications(gomock.Any(), &input).Return(&support.DescribeCommunicationsOutput{Communications: communications}, nil)
	return nil
}
