package route53

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func HealthChecks() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_health_checks",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheck.html`,
		Resolver:    fetchRoute53HealthChecks,
		Transform:   transformers.TransformWithStruct(&Route53HealthCheckWrapper{}, transformers.WithUnwrapStructFields("HealthCheck")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveHealthCheckArn(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Description: `The tags associated with the health check.`,
			},
			{
				Name:     "cloud_watch_alarm_configuration_dimensions",
				Type:     schema.TypeJSON,
				Resolver: resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions,
			},
		},
	}
}
