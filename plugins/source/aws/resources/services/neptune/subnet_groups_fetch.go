package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchNeptuneSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := neptune.DescribeDBSubnetGroupsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}

	c := meta.(*client.Client)
	svc := c.Services().Neptune
	for {
		response, err := svc.DescribeDBSubnetGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DBSubnetGroups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveNeptuneSubnetGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.DBSubnetGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Neptune
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: s.DBSubnetGroupArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
