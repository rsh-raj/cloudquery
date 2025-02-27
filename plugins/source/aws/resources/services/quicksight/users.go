package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "aws_quicksight_users",
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_User.html",
		Resolver:    fetchQuicksightUsers,
		Transform:   transformers.TransformWithStruct(&types.User{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}
