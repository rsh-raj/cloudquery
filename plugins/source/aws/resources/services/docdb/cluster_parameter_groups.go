package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ClusterParameterGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_cluster_parameter_groups",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html`,
		Resolver:    fetchDocdbClusterParameterGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:   transformers.TransformWithStruct(&types.DBClusterParameterGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDBClusterParameterGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: resolveDocdbClusterParameterGroupParameters,
			},
			{
				Name:     "db_cluster_parameter_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroupName"),
			},
			{
				Name:     "db_parameter_group_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupFamily"),
			},
		},
	}
}
