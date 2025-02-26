package account

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchAccountAlternateContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Account
	var contactTypes types.AlternateContactType
	for _, acType := range contactTypes.Values() {
		var input account.GetAlternateContactInput
		input.AlternateContactType = acType
		output, err := svc.GetAlternateContact(ctx, &input)
		if err != nil {
			if client.IsAWSError(err, "ResourceNotFoundException") {
				continue
			}
			return err
		}
		res <- output.AlternateContact
	}
	return nil
}
