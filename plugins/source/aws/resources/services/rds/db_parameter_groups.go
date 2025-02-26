package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func DbParameterGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_db_parameter_groups",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBParameterGroup.html`,
		Resolver:    fetchRdsDbParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.DBParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBParameterGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsDbParameterGroupTags,
			},
		},

		Relations: []*schema.Table{
			DbParameterGroupDbParameters(),
		},
	}
}
