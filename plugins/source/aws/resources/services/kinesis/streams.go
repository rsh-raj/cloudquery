package kinesis

import (
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Streams() *schema.Table {
	return &schema.Table{
		Name:                "aws_kinesis_streams",
		Description:         `https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html`,
		Resolver:            fetchKinesisStreams,
		PreResourceResolver: getStream,
		Transform:           transformers.TransformWithStruct(&types.StreamDescriptionSummary{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("kinesis"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKinesisStreamTags,
			},
		},
	}
}
