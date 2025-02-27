package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchGlacierDataRetrievalPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier

	response, err := svc.GetDataRetrievalPolicy(ctx, &glacier.GetDataRetrievalPolicyInput{})
	if err != nil {
		return err
	}
	res <- response.Policy
	return nil
}
