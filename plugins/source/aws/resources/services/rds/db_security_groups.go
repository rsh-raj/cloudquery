package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func DbSecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_db_security_groups",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSecurityGroup.html`,
		Resolver:    fetchRdsDbSecurityGroups,
		Transform: transformers.TransformWithStruct(
			&types.DBSecurityGroup{},
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"e_c2": "ec2"})),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer("rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSecurityGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRdsDbSecurityGroupTags,
			},
		},
	}
}
