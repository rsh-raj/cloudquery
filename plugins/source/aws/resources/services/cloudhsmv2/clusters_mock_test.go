package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildHSMClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudhsmv2Client(ctrl)

	var clusters []types.Cluster
	if err := faker.FakeObject(&clusters); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeClusters(
		gomock.Any(),
		&cloudhsmv2.DescribeClustersInput{},
		gomock.Any(),
	).Return(
		&cloudhsmv2.DescribeClustersOutput{Clusters: clusters},
		nil,
	)

	return client.Services{Cloudhsmv2: mock}
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildHSMClusters, client.TestOptions{})
}
