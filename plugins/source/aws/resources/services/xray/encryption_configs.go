package xray

import (
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func EncryptionConfigs() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_encryption_configs",
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_EncryptionConfig.html`,
		Resolver:    fetchXrayEncryptionConfigs,
		Transform:   transformers.TransformWithStruct(&types.EncryptionConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
