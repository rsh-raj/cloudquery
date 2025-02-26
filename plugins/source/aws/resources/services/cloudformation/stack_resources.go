package cloudformation

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func stackResources() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudformation_stack_resources",
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html`,
		Resolver:    fetchCloudformationStackResources,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudformation"),
		Transform:   transformers.TransformWithStruct(&types.StackResourceSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "stack_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
