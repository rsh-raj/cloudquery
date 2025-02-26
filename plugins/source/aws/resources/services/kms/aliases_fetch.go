package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchKmsAliases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input kms.ListAliasesInput
	c := meta.(*client.Client)
	svc := c.Services().Kms
	paginator := kms.NewListAliasesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Aliases
	}
	return nil
}
