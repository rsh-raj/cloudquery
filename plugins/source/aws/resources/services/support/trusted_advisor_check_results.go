package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func trustedAdvisorCheckResults() *schema.Table {
	return &schema.Table{
		Name:        "aws_support_trusted_advisor_check_results",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckResult.html`,
		Resolver:    fetchTrustedAdvisorCheckResults,
		Transform:   transformers.TransformWithStruct(&types.TrustedAdvisorCheckResult{}, transformers.WithPrimaryKeys("CheckId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
	}
}

func fetchTrustedAdvisorCheckResults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	// No need to get the result for each language, as those are the same have the same check id
	if c.LanguageCode != "en" {
		return nil
	}
	svc := c.Services().Support
	check := parent.Item.(types.TrustedAdvisorCheckDescription)
	input := support.DescribeTrustedAdvisorCheckResultInput{CheckId: check.Id}

	response, err := svc.DescribeTrustedAdvisorCheckResult(ctx, &input)
	if err != nil {
		return err
	}

	res <- response.Result

	return nil
}

func mockCheckResults(check types.TrustedAdvisorCheckDescription, m *mocks.MockSupportClient) error {
	result := types.TrustedAdvisorCheckResult{}
	err := faker.FakeObject(&result)
	if err != nil {
		return err
	}

	input := support.DescribeTrustedAdvisorCheckResultInput{CheckId: check.Id}
	m.EXPECT().DescribeTrustedAdvisorCheckResult(gomock.Any(), &input).Return(&support.DescribeTrustedAdvisorCheckResultOutput{Result: &result}, nil)
	return nil
}
