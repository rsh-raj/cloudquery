package waf

import (
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Rules() *schema.Table {
	return &schema.Table{
		Name:        "aws_waf_rules",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_RuleSummary.html`,
		Resolver:    fetchWafRules,
		Transform:   transformers.TransformWithStruct(&types.Rule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafRuleArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafRuleTags,
			},
		},
	}
}
