package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func UsageReportSubscriptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_usage_report_subscriptions",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UsageReportSubscription.html`,
		Resolver:    fetchAppstreamUsageReportSubscriptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Transform:   transformers.TransformWithStruct(&types.UsageReportSubscription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "s3_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("S3BucketName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
