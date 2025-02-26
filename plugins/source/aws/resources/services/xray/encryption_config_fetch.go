package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchXrayEncryptionConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Xray
	input := xray.GetEncryptionConfigInput{}
	output, err := svc.GetEncryptionConfig(ctx, &input)
	if err != nil {
		return err
	}
	res <- output.EncryptionConfig
	return nil
}
