package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/iam/models"
)

func fetchIamSamlIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Iam
	response, err := svc.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
	if err != nil {
		return err
	}

	res <- response.SAMLProviderList
	return nil
}

func getSamlIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Iam
	p := resource.Item.(types.SAMLProviderListEntry)

	providerResponse, err := svc.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{SAMLProviderArn: p.Arn})
	if err != nil {
		return err
	}

	resource.Item = models.IAMSAMLIdentityProviderWrapper{
		GetSAMLProviderOutput: providerResponse,
		Arn:                   *p.Arn,
		Tags:                  client.TagsToMap(providerResponse.Tags),
	}
	return nil
}
