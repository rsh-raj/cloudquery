package frauddetector

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func ModelVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_model_versions",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_ModelVersionDetail.html`,
		Resolver:    fetchFrauddetectorModelVersions,
		Transform:   transformers.TransformWithStruct(&types.ModelVersionDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
