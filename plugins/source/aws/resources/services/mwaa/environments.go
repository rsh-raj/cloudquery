package mwaa

import (
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Environments() *schema.Table {
	return &schema.Table{
		Name:                "aws_mwaa_environments",
		Description:         `https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html`,
		Resolver:            fetchMwaaEnvironments,
		Transform:           transformers.TransformWithStruct(&types.Environment{}),
		PreResourceResolver: getEnvironment,
		Multiplex:           client.ServiceAccountRegionMultiplexer("airflow"),
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
