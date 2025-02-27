package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/ec2/models"
)

func fetchEc2RegionalConfigs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	svc := c.Services().Ec2
	var regionalConfig models.RegionalConfig
	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{})
	if err != nil {
		return err
	}
	regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
	if err != nil {
		return err
	}

	if ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}
	res <- regionalConfig
	return nil
}
