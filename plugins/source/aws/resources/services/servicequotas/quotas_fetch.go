package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchServicequotasQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Servicequotas
	service := parent.Item.(types.ServiceInfo)
	config := servicequotas.ListServiceQuotasInput{
		ServiceCode: service.ServiceCode,
		MaxResults:  defaultMaxResults,
	}
	quotasPaginator := servicequotas.NewListServiceQuotasPaginator(svc, &config)
	for quotasPaginator.HasMorePages() {
		output, err := quotasPaginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Quotas
	}
	return nil
}
