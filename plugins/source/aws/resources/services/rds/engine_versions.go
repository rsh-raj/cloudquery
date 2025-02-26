package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func EngineVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_engine_versions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html`,
		Resolver:    fetchRdsEngineVersions,
		Transform:   transformers.TransformWithStruct(&types.DBEngineVersion{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "_engine_version_hash",
				Type:     schema.TypeString,
				Resolver: client.ResolveObjectHash,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tag_list",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
		},

		Relations: []*schema.Table{
			ClusterParameters(),
		},
	}
}
