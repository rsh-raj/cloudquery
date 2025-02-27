package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Operations() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_operations",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html`,
		Resolver:    fetchApprunnerOperations,
		Transform:   transformers.TransformWithStruct(&types.OperationSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
