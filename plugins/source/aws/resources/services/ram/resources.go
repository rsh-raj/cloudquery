package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Resources() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resources",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html`,
		Resolver:    fetchRamResources,
		Transform:   transformers.TransformWithStruct(&types.Resource{}, transformers.WithPrimaryKeys("Arn", "ResourceShareArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
