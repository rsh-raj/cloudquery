package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func DataCatalogs() *schema.Table {
	return &schema.Table{
		Name:                "aws_athena_data_catalogs",
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html`,
		Resolver:            fetchAthenaDataCatalogs,
		PreResourceResolver: getDataCatalog,
		Multiplex:           client.ServiceAccountRegionMultiplexer("athena"),
		Transform:           transformers.TransformWithStruct(&types.DataCatalog{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveAthenaDataCatalogArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveAthenaDataCatalogTags,
			},
		},

		Relations: []*schema.Table{
			DataCatalogDatabases(),
		},
	}
}
