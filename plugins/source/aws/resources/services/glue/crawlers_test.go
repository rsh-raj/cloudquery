package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildCrawlers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var crawler glue.GetCrawlersOutput
	if err := faker.FakeObject(&crawler); err != nil {
		t.Fatal(err)
	}
	crawler.NextToken = nil
	m.EXPECT().GetCrawlers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&crawler, nil)

	var tags glue.GetTagsOutput
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{
		Glue: m,
	}
}

func TestCrawlers(t *testing.T) {
	client.AwsMockTestHelper(t, Crawlers(), buildCrawlers, client.TestOptions{})
}
