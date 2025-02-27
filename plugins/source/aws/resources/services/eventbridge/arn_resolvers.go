package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func resolveArchiveArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "events",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "archive/" + aws.ToString(resource.Item.(types.Archive).ArchiveName),
	}

	return resource.Set(c.Name, a.String())
}

func resolveReplayArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "events",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "replay/" + aws.ToString(resource.Item.(types.Replay).ReplayName),
	}

	return resource.Set(c.Name, a.String())
}
