package wafregional

import (
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func RateBasedRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafregional_rate_based_rules",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RateBasedRule.html`,
		Resolver:    fetchWafregionalRateBasedRules,
		Transform:   transformers.TransformWithStruct(&types.RateBasedRule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafregionalRateBasedRuleArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafregionalRateBasedRuleTags,
			},
		},
	}
}
