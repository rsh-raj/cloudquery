package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchClusterInstanceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(*types.Cluster)
	// instance fleets and instance groups are mutually exclusive
	if cluster.InstanceCollectionType != types.InstanceCollectionTypeInstanceGroup {
		return nil
	}
	config := emr.ListInstanceGroupsInput{
		ClusterId: cluster.Id,
	}
	c := meta.(*client.Client)
	svc := c.Services().Emr
	for {
		response, err := svc.ListInstanceGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.InstanceGroups

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
