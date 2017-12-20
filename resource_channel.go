package main

import (
	"log"

	"github.com/nlopes/slack"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Exists: resourceChannelExists,

		Schema: map[string]*schema.Schema{
			"channel_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of Slack Channel that will be created",
				Required:    true,
			},
		},
	}
}

func resourceChannelExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("[INFO] Creating Slack channel")

	channelName := d.Get("channel_name").(string)
	api := slack.New(meta.(*Config).APIToken)

	api.SetDebug(true)

	channel, err := api.CreateChannel(channelName)

	if err != nil {
		log.Println("[DEBUG] Cannot create Slack chanel: [", channelName, "]")
		return nil
	}

	// ResourceID is set as Slack Channel ID
	d.SetId(channel.ID)
	return nil

}

func resourceChannelRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
