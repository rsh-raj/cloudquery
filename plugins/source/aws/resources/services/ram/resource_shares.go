package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ResourceShares() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_shares",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShare.html`,
		Resolver:    fetchRamResourceShares,
		Transform:   transformers.TransformWithStruct(&types.ResourceShare{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceShareArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			ResourceSharePermissions(),
		},
	}
}
