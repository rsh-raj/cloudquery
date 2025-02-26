package amp

import (
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Workspaces() *schema.Table {
	return &schema.Table{
		Name:                "aws_amp_workspaces",
		Description:         `https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference.html#AMP-APIReference-WorkspaceDescription`,
		Resolver:            fetchAmpWorkspaces,
		PreResourceResolver: describeWorkspace,
		Multiplex:           client.ServiceAccountRegionMultiplexer("amp"),
		Transform:           transformers.TransformWithStruct(&types.WorkspaceDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "alert_manager_definition",
				Type:     schema.TypeJSON,
				Resolver: describeAlertManagerDefinition,
			},
			{
				Name:     "logging_configuration",
				Type:     schema.TypeJSON,
				Resolver: describeLoggingConfiguration,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			RuleGroupsNamespaces(),
		},
	}
}
