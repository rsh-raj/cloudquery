package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func fetchFrauddetectorDetectors(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	paginator := frauddetector.NewGetDetectorsPaginator(meta.(*client.Client).Services().Frauddetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Detectors
	}
	return nil
}
