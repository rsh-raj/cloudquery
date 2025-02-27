package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "aws_identitystore_users",
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html`,
		Resolver:    fetchIdentitystoreUsers,
		Transform:   transformers.TransformWithStruct(&types.User{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("identitystore"),
	}
}
