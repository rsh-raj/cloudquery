package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func PendingMaintenanceActions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_pending_maintenance_actions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_PendingMaintenanceAction.html`,
		Resolver:    fetchDocdbPendingMaintenanceActions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:   transformers.TransformWithStruct(&types.ResourcePendingMaintenanceActions{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
