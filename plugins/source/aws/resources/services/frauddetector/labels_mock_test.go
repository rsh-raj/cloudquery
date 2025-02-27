package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildLabels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.Label{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetLabels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetLabelsOutput{Labels: []types.Label{data}}, nil,
	)

	addTagsCall(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestLabels(t *testing.T) {
	client.AwsMockTestHelper(t, Labels(), buildLabels, client.TestOptions{})
}
