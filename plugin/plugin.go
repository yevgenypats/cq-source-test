package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yevgenypats/cq-source-test/client"
	"github.com/yevgenypats/cq-source-test/tables"
)

const exampleConfig = `
# This is an example config file for the plugin.
account_ids: []
`

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"test",
		Version,
		[]*schema.Table{
			tables.TestSomeTable(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
