package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_buckets",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Bucket.html`,
		Resolver:    fetchLightsailBuckets,
		Transform:   transformers.TransformWithStruct(&types.Bucket{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "able_to_update_bundle",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AbleToUpdateBundle"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			BucketAccessKeys(),
		},
	}
}
