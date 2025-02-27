package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_identitystore_groups",
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html`,
		Resolver:    fetchIdentitystoreGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("identitystore"),

		Relations: []*schema.Table{
			GroupMemberships(),
		},
	}
}
