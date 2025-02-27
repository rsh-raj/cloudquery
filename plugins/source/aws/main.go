package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/plugin"
)

const sentryDSN = "https://6c6b72bc946844cb8471f49eba485cde@o1396617.ingest.sentry.io/6747636"

func main() {
	serve.Source(plugin.AWS(), serve.WithSourceSentryDSN(sentryDSN))
}
