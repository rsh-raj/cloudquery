package workspaces

import (
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Workspaces() *schema.Table {
	return &schema.Table{
		Name:        "aws_workspaces_workspaces",
		Description: `https://docs.aws.amazon.com/workspaces/latest/api/API_Workspace.html`,
		Resolver:    fetchWorkspacesWorkspaces,
		Transform:   transformers.TransformWithStruct(&types.Workspace{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("workspaces"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWorkspaceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
