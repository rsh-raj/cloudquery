package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/iam/models"
)

func SamlIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_saml_identity_providers",
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html`,
		Resolver:            fetchIamSamlIdentityProviders,
		PreResourceResolver: getSamlIdentityProvider,
		Transform: transformers.TransformWithStruct(
			&models.IAMSAMLIdentityProviderWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithSkipFields("ResultMetadata"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
