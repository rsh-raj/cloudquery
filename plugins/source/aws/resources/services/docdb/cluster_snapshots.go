package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ClusterSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_cluster_snapshots",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshot.html`,
		Resolver:    fetchDocdbClusterSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:   transformers.TransformWithStruct(&types.DBClusterSnapshot{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDBClusterSnapshotTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterSnapshotArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: resolveDocdbClusterSnapshotAttributes,
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_cluster_snapshot_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterSnapshotIdentifier"),
			},
		},
	}
}
