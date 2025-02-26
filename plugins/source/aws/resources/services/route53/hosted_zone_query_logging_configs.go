package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func HostedZoneQueryLoggingConfigs() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_hosted_zone_query_logging_configs",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_QueryLoggingConfig.html`,
		Resolver:    fetchRoute53HostedZoneQueryLoggingConfigs,
		Transform:   transformers.TransformWithStruct(&types.QueryLoggingConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveRoute53HostedZoneQueryLoggingConfigsArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "hosted_zone_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
