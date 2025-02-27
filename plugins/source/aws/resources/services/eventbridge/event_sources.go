package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func EventSources() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_event_sources",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventSource.html`,
		Resolver:    fetchEventbridgeEventSources,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
		Transform:   transformers.TransformWithStruct(&types.EventSource{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
