package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func BucketAccessKeys() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_bucket_access_keys",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_AccessKey.html`,
		Resolver:    fetchLightsailBucketAccessKeys,
		Transform:   transformers.TransformWithStruct(&types.AccessKey{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
