package appsync

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func GraphqlApis() *schema.Table {
	return &schema.Table{
		Name:        "aws_appsync_graphql_apis",
		Description: `https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html`,
		Resolver:    fetchAppsyncGraphqlApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appsync"),
		Transform:   transformers.TransformWithStruct(&types.GraphqlApi{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
