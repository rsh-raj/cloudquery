package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func InstanceSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_instance_snapshots",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstanceSnapshot.html`,
		Resolver:    fetchLightsailInstanceSnapshots,
		Transform:   transformers.TransformWithStruct(&types.InstanceSnapshot{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name: "arn",
				Type: schema.TypeString,
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
	}
}
