package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/TimDurward/terraform-provider-slack/slack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: slack.Provider,
	})
}
