package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider to manage slack.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
				Description: "Slack Authentication Token for api.slack.com",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"slack_channel": resourceChannel(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"slack_user": dataSourceSlackUser(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		APIToken: d.Get("api_token").(string),
	}
	return config, nil
}
