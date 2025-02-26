package ses

import (
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func CustomVerificationEmailTemplates() *schema.Table {
	return &schema.Table{
		Name:                "aws_ses_custom_verification_email_templates",
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetCustomVerificationEmailTemplate.html`,
		Resolver:            fetchSesCustomVerificationEmailTemplates,
		PreResourceResolver: getCustomVerificationEmailTemplate,
		Transform: transformers.TransformWithStruct(
			&sesv2.GetCustomVerificationEmailTemplateOutput{},
			transformers.WithSkipFields("ResultMetadata"),
			transformers.WithNameTransformer(client.CreateTrimPrefixTransformer("template_")),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCustomVerificationEmailTemplateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
