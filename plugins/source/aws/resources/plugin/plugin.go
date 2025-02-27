package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

var (
	Version = "Development"
)

func AWS() *source.Plugin {
	return source.NewPlugin(
		"aws",
		Version,
		tables(),
		client.Configure,
	)
}
