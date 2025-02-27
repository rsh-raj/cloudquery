package elbv1

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/elbv1/models"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv1_load_balancers",
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html`,
		Resolver:    fetchElbv1LoadBalancers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&models.ELBv1LoadBalancerWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveLoadBalancerARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			LoadBalancerPolicies(),
		},
	}
}
