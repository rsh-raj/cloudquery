package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_snapshots",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchElasticacheSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Transform:   transformers.TransformWithStruct(&types.Snapshot{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
