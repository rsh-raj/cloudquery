package dax

import (
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_dax_clusters",
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html`,
		Resolver:    fetchDaxClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dax"),
		Transform:   transformers.TransformWithStruct(&types.Cluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveClusterTags,
			},
		},
	}
}
