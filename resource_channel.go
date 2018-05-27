package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/timdurward/slack"
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

			"channel_topic": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Sets the topic for a channel",
				Optional:    true,
			},
		},
	}
}

func resourceChannelExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	// Create Slack Channel
	//TODO: Add 21 Character Limit check
	channel, err := api.CreateChannel(d.Get("channel_name").(string))
	if err != nil {
		return err
	}
	d.SetId(channel.ID)

	// Create Slack Channel Topic
	if _, err := api.SetChannelTopic(channel.ID, d.Get("channel_topic").(string)); err != nil {
		return err
	}

	return nil
}

func resourceChannelRead(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	// Checks if Slack Channel exists, if not remove resource from state
	_, err := api.GetChannelInfo(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	// Deletes Slack Channel and clears state
	_, err := api.DeleteChannel(d.Id())
	if err != nil {
		return err
	}

	return nil
}
