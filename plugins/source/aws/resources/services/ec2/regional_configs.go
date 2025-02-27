package ec2

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/ec2/models"
)

func RegionalConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_regional_configs",
		Resolver:  fetchEc2RegionalConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
		Transform: transformers.TransformWithStruct(&models.RegionalConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
