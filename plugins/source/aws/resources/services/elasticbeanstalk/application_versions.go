package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ApplicationVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_application_versions",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html`,
		Resolver:    fetchElasticbeanstalkApplicationVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.ApplicationVersionDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationVersionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
