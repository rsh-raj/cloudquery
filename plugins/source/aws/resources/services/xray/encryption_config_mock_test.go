package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildEncryptionConfig(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	var config types.EncryptionConfig
	if err := faker.FakeObject(&config); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetEncryptionConfig(
		gomock.Any(),
		&xray.GetEncryptionConfigInput{},
		gomock.Any(),
	).Return(
		&xray.GetEncryptionConfigOutput{
			EncryptionConfig: &config,
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestXrayEncryptionConfig(t *testing.T) {
	client.AwsMockTestHelper(t, EncryptionConfigs(), buildEncryptionConfig, client.TestOptions{})
}
