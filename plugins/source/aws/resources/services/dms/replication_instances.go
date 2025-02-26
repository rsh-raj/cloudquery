package dms

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/dms/models"
)

func ReplicationInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_dms_replication_instances",
		Description: `https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html`,
		Resolver:    fetchDmsReplicationInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dms"),
		Transform:   transformers.TransformWithStruct(&models.ReplicationInstanceWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
