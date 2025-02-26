package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchRoute53Domains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53domains
	var input route53domains.ListDomainsInput

	for {
		output, err := svc.ListDomains(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.Domains

		if aws.ToString(output.NextPageMarker) == "" {
			break
		}
		input.Marker = output.NextPageMarker
	}
	return nil
}
func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53domains
	v := resource.Item.(types.DomainSummary)

	d, err := svc.GetDomainDetail(ctx, &route53domains.GetDomainDetailInput{DomainName: v.DomainName})
	if err != nil {
		return err
	}

	resource.Item = d

	return resource.Set("transfer_lock", aws.ToBool(v.TransferLock))
}

func resolveRoute53DomainTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)
	svc := c.Services().Route53domains
	d := resource.Item.(*route53domains.GetDomainDetailOutput)
	out, err := svc.ListTagsForDomain(ctx, &route53domains.ListTagsForDomainInput{DomainName: d.DomainName})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, client.TagsToMap(out.TagList))
}
