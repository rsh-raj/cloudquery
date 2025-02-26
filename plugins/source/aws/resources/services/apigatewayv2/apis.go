package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Apis() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_apis",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Api.html`,
		Resolver:    fetchApigatewayv2Apis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&types.Api{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiId"),
			},
		},
		Relations: []*schema.Table{
			ApiAuthorizers(),
			ApiDeployments(),
			ApiIntegrations(),
			ApiModels(),
			ApiRoutes(),
			ApiStages(),
		},
	}
}
