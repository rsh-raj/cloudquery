package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func UserGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_user_groups",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html`,
		Resolver:    fetchIamUserGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
