package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchRdsSubnetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config rds.DescribeDBSubnetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Rds
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
