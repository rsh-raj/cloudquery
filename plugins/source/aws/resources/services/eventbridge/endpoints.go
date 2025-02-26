package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Endpoints() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_endpoints",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Endpoint.html`,
		Resolver:    fetchEventbridgeEndpoints,
		Multiplex:   client.AccountMultiplex,
		Transform:   transformers.TransformWithStruct(&types.Endpoint{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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
