package main

import (
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

			"is_private": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Boolean value to check if Slack channel will be private or not",
				Required:    true,
			},
		},
	}
}

func resourceChannelExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	channelName := d.Get("channel_name").(string)
	isPrivate := d.Get("is_private").(bool)

	api := slack.New(meta.(*Config).APIToken)

	if isPrivate {
		channel, err := api.CreateGroup(channelName)

		if err != nil {
			return err
		}

		d.SetId(channel.ID)
		return nil
	}

	channel, err := api.CreateChannel(channelName)
	if err != nil {
		return err
	}
	d.SetId(channel.ID)
	return nil
}

func resourceChannelRead(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	_, channelResponse := api.GetChannelInfo(d.Id())

	if channelResponse != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
