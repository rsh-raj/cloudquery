package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchLightsailInstanceSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetInstanceSnapshotsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetInstanceSnapshots(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.InstanceSnapshots
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
