package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func DataRetrievalPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_glacier_data_retrieval_policies",
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html`,
		Resolver:    fetchGlacierDataRetrievalPolicies,
		Transform:   transformers.TransformWithStruct(&types.DataRetrievalPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glacier"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
