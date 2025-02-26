package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchKafkaClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input kafka.ListClustersV2Input
	c := meta.(*client.Client)
	svc := c.Services().Kafka
	for {
		response, err := svc.ListClustersV2(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ClusterInfoList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kafka
	var input kafka.DescribeClusterV2Input = describeClustersInput(resource)
	output, err := svc.DescribeClusterV2(ctx, &input)
	if err != nil {
		return err
	}
	resource.Item = output.ClusterInfo
	return nil
}
