package cloudtrail

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/cloudtrail/models"
)

func Trails() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudtrail_trails",
		Description: `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Trail.html`,
		Resolver:    fetchCloudtrailTrails,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudtrail"),
		Transform:   transformers.TransformWithStruct(&models.CloudTrailWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "cloudwatch_logs_log_group_name",
				Type:     schema.TypeString,
				Resolver: resolveCloudtrailTrailCloudwatchLogsLogGroupName,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrailARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudTrailStatus,
			},
		},
		Relations: []*schema.Table{
			TrailEventSelectors(),
		},
	}
}
