package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func DirectoryConfigs() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_directory_configs",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html`,
		Resolver:    fetchAppstreamDirectoryConfigs,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Transform:   transformers.TransformWithStruct(&types.DirectoryConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "directory_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
