package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ClusterParameters() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_cluster_parameters",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Parameter.html`,
		Resolver:    fetchDocdbClusterParameters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
