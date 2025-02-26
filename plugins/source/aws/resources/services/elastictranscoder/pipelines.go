package elastictranscoder

import (
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Pipelines() *schema.Table {
	return &schema.Table{
		Name:        "aws_elastictranscoder_pipelines",
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-pipelines.html`,
		Resolver:    fetchElastictranscoderPipelines,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elastictranscoder"),
		Transform:   transformers.TransformWithStruct(&types.Pipeline{}),
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

		Relations: []*schema.Table{
			PipelineJobs(),
		},
	}
}
