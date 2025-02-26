package account

import (
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func AlternateContacts() *schema.Table {
	return &schema.Table{
		Name:        "aws_account_alternate_contacts",
		Description: `https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html`,
		Resolver:    fetchAccountAlternateContacts,
		Multiplex:   client.ServiceAccountRegionMultiplexer("account"),
		Transform:   transformers.TransformWithStruct(&types.AlternateContact{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "alternate_contact_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlternateContactType"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
