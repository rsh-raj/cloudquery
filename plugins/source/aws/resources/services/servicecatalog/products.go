package servicecatalog

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicecatalog_products",
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html`,
		Resolver:    fetchServicecatalogProducts,
		Transform:   transformers.TransformWithStruct(&types.ProductViewDetail{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("servicecatalog"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveProductTags,
			},
		},
	}
}
