package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildAppVersionResourceMappings(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppVersionResourceMappingsOutput
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.NextToken = nil
	mock.EXPECT().ListAppVersionResourceMappings(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
}
