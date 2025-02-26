package glacier

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
	"github.com/stretchr/testify/require"
)

func buildDRPMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlacierClient(ctrl)

	p := glacier.GetDataRetrievalPolicyOutput{}
	require.NoError(t, faker.FakeObject(&p))
	m.EXPECT().GetDataRetrievalPolicy(gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		Glacier: m,
	}
}

func TestDataRetrievalPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, DataRetrievalPolicies(), buildDRPMock, client.TestOptions{})
}
