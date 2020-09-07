package main

import (
	"github.com/TimDurward/terraform-provider-slack/slack"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: slack.Provider,
	})
}
