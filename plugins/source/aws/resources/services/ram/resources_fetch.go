package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchRamResources(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	err := fetchRamResourcesByOwner(ctx, meta, types.ResourceOwnerSelf, res)
	if err != nil {
		return err
	}
	err = fetchRamResourcesByOwner(ctx, meta, types.ResourceOwnerOtherAccounts, res)
	if err != nil {
		return err
	}
	return nil
}

func fetchRamResourcesByOwner(ctx context.Context, meta schema.ClientMeta, shareType types.ResourceOwner, res chan<- any) error {
	input := &ram.ListResourcesInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: shareType,
	}
	paginator := ram.NewListResourcesPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Resources
	}
	return nil
}
