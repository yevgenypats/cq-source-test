package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/yevgenypats/cq-source-test/plugin"
)

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
	})
}
