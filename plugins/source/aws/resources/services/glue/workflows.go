package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:                "aws_glue_workflows",
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_Workflow.html`,
		Resolver:            fetchGlueWorkflows,
		PreResourceResolver: getWorkflow,
		Transform:           transformers.TransformWithStruct(&types.Workflow{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueWorkflowArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueWorkflowTags,
			},
		},
	}
}
