package slack

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/slack-go/slack"
)

func dataSourceSlackUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"real_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceSlackUserRead(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	email := d.Get("email").(string)
	log.Printf("[INFO] Reading Slack user '%s'", email)

	users, err := api.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Profile.Email == email {
			log.Printf("[DEBUG] Slack user: %v", user)
			d.SetId(user.ID)
			d.Set("name", user.Name)
			d.Set("real_name", user.RealName)
			d.Set("email", user.Profile.Email)
			return nil
		}
	}

	return fmt.Errorf("Invalid user with email '%s'", email)
}
