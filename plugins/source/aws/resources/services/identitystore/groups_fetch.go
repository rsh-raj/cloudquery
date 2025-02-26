package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchIdentitystoreGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	svc := meta.(*client.Client).Services().Identitystore
	config := identitystore.ListGroupsInput{}
	config.IdentityStoreId = instance.IdentityStoreId
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Groups

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
