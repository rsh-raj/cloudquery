package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ReservedCacheNodes() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_reserved_cache_nodes",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNode.html`,
		Resolver:    fetchElasticacheReservedCacheNodes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Transform:   transformers.TransformWithStruct(&types.ReservedCacheNode{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReservationARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
