package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildSeverityLevels(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	levels := []types.SeverityLevel{}
	err := faker.FakeObject(&levels)
	if err != nil {
		t.Fatal(err)
	}

	for _, languageCode := range severitySupportedLanguageCodes {
		m.EXPECT().DescribeSeverityLevels(gomock.Any(), &support.DescribeSeverityLevelsInput{Language: aws.String(languageCode)}).Return(&support.DescribeSeverityLevelsOutput{SeverityLevels: levels}, nil)
	}

	return client.Services{
		Support: m,
	}
}

func TestSeverityLevels(t *testing.T) {
	client.AwsMockTestHelper(t, SeverityLevels(), buildSeverityLevels, client.TestOptions{})
}
