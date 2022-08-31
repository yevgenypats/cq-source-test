package plugin

import (
	"github.com/cloudquery/cq-source-test/client"
	"github.com/cloudquery/cq-source-test/tables"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
# This is an example config file for the plugin.
account_ids: []
`

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return &plugins.SourcePlugin{
		Name:      "test",
		Version:   Version,
		Configure: client.Configure,
		Tables: []*schema.Table{
			tables.TestSomeTable(),
		},
		ExampleConfig: exampleConfig,
	}
}
