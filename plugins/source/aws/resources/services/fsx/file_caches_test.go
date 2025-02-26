package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
	"github.com/stretchr/testify/require"
)

func buildFileCachesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var fc types.FileCache
	require.NoError(t, faker.FakeObject(&fc))
	m.EXPECT().DescribeFileCaches(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&fsx.DescribeFileCachesOutput{FileCaches: []types.FileCache{fc}},
		nil,
	)

	var tags []types.Tag
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&fsx.ListTagsForResourceOutput{Tags: tags},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestFileCaches(t *testing.T) {
	client.AwsMockTestHelper(t, FileCaches(), buildFileCachesMock, client.TestOptions{})
}
