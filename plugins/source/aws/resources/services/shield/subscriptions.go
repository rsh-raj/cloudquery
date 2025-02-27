package shield

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_shield_subscriptions",
		Description: `https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Subscription.html`,
		Resolver:    fetchShieldSubscriptions,
		Transform:   transformers.TransformWithStruct(&types.Subscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("shield"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
