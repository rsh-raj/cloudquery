package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func DiskSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_disk_snapshots",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DiskSnapshot.html`,
		Resolver:    fetchLightsailDiskSnapshots,
		Transform:   transformers.TransformWithStruct(&types.DiskSnapshot{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "disk_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
