package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/iam/models"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:      "aws_iam_accounts",
		Resolver:  fetchIamAccounts,
		Transform: transformers.TransformWithStruct(&models.Account{}),
		Multiplex: client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}
