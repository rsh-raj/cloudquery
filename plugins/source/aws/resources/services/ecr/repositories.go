package ecr

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repositories",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_Repository.html`,
		Resolver:    fetchEcrRepositories,
		Multiplex:   client.ServiceAccountRegionMultiplexer("api.ecr"),
		Transform:   transformers.TransformWithStruct(&types.Repository{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRepositoryTags,
			},
			{
				Name:     "policy_text",
				Type:     schema.TypeJSON,
				Resolver: resolveRepositoryPolicy,
			},
		},

		Relations: []*schema.Table{
			RepositoryImages(),
		},
	}
}
