package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Streams() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_streams",
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_StreamInfo.html`,
		Resolver:    fetchIotStreams,
		Transform:   transformers.TransformWithStruct(&types.StreamInfo{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
