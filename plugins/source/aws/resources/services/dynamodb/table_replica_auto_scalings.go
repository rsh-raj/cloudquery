package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func TableReplicaAutoScalings() *schema.Table {
	return &schema.Table{
		Name:        "aws_dynamodb_table_replica_auto_scalings",
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html`,
		Resolver:    fetchDynamodbTableReplicaAutoScalings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dynamodb"),
		Transform:   transformers.TransformWithStruct(&types.ReplicaAutoScalingDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "table_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
