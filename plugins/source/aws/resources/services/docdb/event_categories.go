package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func EventCategories() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_event_categories",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventCategoriesMap.html`,
		Resolver:    fetchDocdbEventCategories,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:   transformers.TransformWithStruct(&types.EventCategoriesMap{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "event_categories",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EventCategories"),
			},
			{
				Name:     "source_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceType"),
			},
		},
	}
}
