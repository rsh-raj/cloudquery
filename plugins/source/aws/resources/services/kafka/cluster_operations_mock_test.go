package kafka

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildKafkaClusterOperationsMock(t *testing.T, m *mocks.MockKafkaClient) {
	object := types.ClusterOperationInfo{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListClusterOperations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kafka.ListClusterOperationsOutput{
			ClusterOperationInfoList: []types.ClusterOperationInfo{object},
		}, nil)
}
