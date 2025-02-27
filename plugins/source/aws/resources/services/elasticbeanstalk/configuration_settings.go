package elasticbeanstalk

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
)

func ConfigurationSettings() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_configuration_settings",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationSettingsDescription.html`,
		Resolver:    fetchElasticbeanstalkConfigurationSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(models.ConfigurationSettingsDescriptionWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "environment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
