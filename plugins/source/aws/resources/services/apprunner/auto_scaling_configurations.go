package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func AutoScalingConfigurations() *schema.Table {
	return &schema.Table{
		Name:                "aws_apprunner_auto_scaling_configurations",
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html`,
		Resolver:            fetchApprunnerAutoScalingConfigurations,
		PreResourceResolver: getAutoScalingConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer("apprunner"),
		Transform:           transformers.TransformWithStruct(&types.AutoScalingConfiguration{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingConfigurationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("AutoScalingConfigurationArn"),
			},
		},
	}
}
