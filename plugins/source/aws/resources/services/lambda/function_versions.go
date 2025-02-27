package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func FunctionVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_function_versions",
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionConfiguration.html`,
		Resolver:    fetchLambdaFunctionVersions,
		Transform:   transformers.TransformWithStruct(&types.FunctionConfiguration{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
