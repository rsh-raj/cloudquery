package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchApprunnerVpcIngressConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListVpcIngressConnectionsInput
	svc := meta.(*client.Client).Services().Apprunner
	paginator := apprunner.NewListVpcIngressConnectionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.VpcIngressConnectionSummaryList
	}
	return nil
}

func getVpcIngressConnection(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Apprunner
	asConfig := resource.Item.(types.VpcIngressConnectionSummary)

	describeTaskDefinitionOutput, err := svc.DescribeVpcIngressConnection(ctx, &apprunner.DescribeVpcIngressConnectionInput{VpcIngressConnectionArn: asConfig.VpcIngressConnectionArn})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput.VpcIngressConnection
	return nil
}
