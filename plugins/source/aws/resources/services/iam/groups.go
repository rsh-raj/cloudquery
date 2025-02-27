package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_groups",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html`,
		Resolver:    fetchIamGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: resolveIamGroupPolicies,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			GroupPolicies(),
		},
	}
}
