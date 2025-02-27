package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchElasticacheUpdateAction(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeUpdateActionsInput

	paginator := elasticache.NewDescribeUpdateActionsPaginator(meta.(*client.Client).Services().Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.UpdateActions
	}
	return nil
}
