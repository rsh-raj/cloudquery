package quicksight

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchQuicksightUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Quicksight
	input := quicksight.ListUsersInput{
		AwsAccountId: aws.String(cl.AccountID),
		Namespace:    aws.String(defaultNamespace),
	}
	var ae smithy.APIError

	for {
		out, err := svc.ListUsers(ctx, &input)
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}

			return err
		}
		res <- out.UserList

		if aws.ToString(out.NextToken) == "" {
			break
		}
		input.NextToken = out.NextToken
	}
	return nil
}
