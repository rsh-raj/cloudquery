package emr

import (
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func BlockPublicAccessConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_emr_block_public_access_configs",
		Resolver:  fetchEmrBlockPublicAccessConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
		Transform: transformers.TransformWithStruct(&emr.GetBlockPublicAccessConfigurationOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
