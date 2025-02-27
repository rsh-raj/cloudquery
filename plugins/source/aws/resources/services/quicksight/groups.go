package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_quicksight_groups",
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Group.html",
		Resolver:    fetchQuicksightGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
		Relations:   []*schema.Table{groupMembers()},
	}
}
