package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func PlanSelections() *schema.Table {
	return &schema.Table{
		Name:        "aws_backup_plan_selections",
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupSelection.html`,
		Resolver:    fetchBackupPlanSelections,
		Multiplex:   client.ServiceAccountRegionMultiplexer("backup"),
		Transform:   transformers.TransformWithStruct(&backup.GetBackupSelectionOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "plan_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
