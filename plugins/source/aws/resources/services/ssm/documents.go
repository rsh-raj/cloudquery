package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Documents() *schema.Table {
	return &schema.Table{
		Name:                "aws_ssm_documents",
		Description:         `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentDescription.html`,
		Resolver:            fetchSsmDocuments,
		PreResourceResolver: getDocument,
		Transform:           transformers.TransformWithStruct(&types.DocumentDescription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveDocumentARN,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "permissions",
				Type:     schema.TypeJSON,
				Resolver: resolveDocumentPermission,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
