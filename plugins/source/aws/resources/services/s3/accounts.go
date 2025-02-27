package s3

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/s3/models"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_accounts",
		Resolver:  fetchS3Accounts,
		Transform: transformers.TransformWithStruct(&models.PublicAccessBlockConfigurationWrapper{}, transformers.WithUnwrapStructFields("PublicAccessBlockConfiguration")),
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
