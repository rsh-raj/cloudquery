package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchShieldProtections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.ListProtectionsInput{}
	for {
		output, err := svc.ListProtections(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.Protections

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveShieldProtectionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Protection)
	cli := meta.(*client.Client)
	svc := cli.Services().Shield
	config := shield.ListTagsForResourceInput{ResourceARN: r.ProtectionArn}

	output, err := svc.ListTagsForResource(ctx, &config, func(o *shield.Options) {
		o.Region = cli.Region
	})
	if err != nil {
		if cli.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
