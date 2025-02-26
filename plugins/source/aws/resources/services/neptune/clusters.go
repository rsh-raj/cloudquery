package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_neptune_clusters",
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-clusters.html#DescribeDBClusters`,
		Resolver:    fetchNeptuneClusters,
		Transform:   transformers.TransformWithStruct(&types.DBCluster{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneClusterTags,
			},
		},
	}
}
