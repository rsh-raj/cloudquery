package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchIdentitystoreGroupMemberships(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Identitystore
	group := parent.Item.(types.Group)
	config := identitystore.ListGroupMembershipsInput{
		GroupId:         group.GroupId,
		IdentityStoreId: group.IdentityStoreId,
	}
	for {
		response, err := svc.ListGroupMemberships(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.GroupMemberships

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
