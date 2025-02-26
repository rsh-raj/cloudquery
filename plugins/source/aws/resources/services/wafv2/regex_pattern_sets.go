package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func RegexPatternSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_wafv2_regex_pattern_sets",
		Description:         `https://docs.aws.amazon.com/waf/latest/APIReference/API_RegexPatternSet.html`,
		Resolver:            fetchWafv2RegexPatternSets,
		PreResourceResolver: getRegexPatternSet,
		Transform:           transformers.TransformWithStruct(&types.RegexPatternSet{}),
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRegexPatternSetTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
