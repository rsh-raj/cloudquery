package cloudfront

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func CachePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudfront_cache_policies",
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html`,
		Resolver:    fetchCloudfrontCachePolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudfront"),
		Transform:   transformers.TransformWithStruct(&types.CachePolicySummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CachePolicy.Id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCachePolicyARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
