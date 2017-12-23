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
				Optional:    true,
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
	isPrivate := d.Get("is_private").(bool)

	if isPrivate {
		err := createPrivateChannel(d, meta)
		if err != nil {
			return err
		}
	} else {
		err := createPublicChannel(d, meta)
		if err != nil {
			return err
		}
	}
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

func createPrivateChannel(d *schema.ResourceData, meta interface{}) error {
	channelName := d.Get("channel_name").(string)
	channelTopic := d.Get("channel_topic").(string)
	api := slack.New(meta.(*Config).APIToken)

	privateChannel, privateChannelerror := api.CreateGroup(channelName)

	if privateChannelerror != nil {
		return privateChannelerror
	}

	_, topicError := api.SetGroupTopic(privateChannel.ID, channelTopic)

	if topicError != nil {
		return topicError
	}

	d.SetId(privateChannel.ID)
	return privateChannelerror
}

func createPublicChannel(d *schema.ResourceData, meta interface{}) error {
	channelName := d.Get("channel_name").(string)
	channelTopic := d.Get("channel_topic").(string)
	api := slack.New(meta.(*Config).APIToken)

	publicChannel, publicChannelError := api.CreateGroup(channelName)

	if publicChannelError != nil {
		return publicChannelError
	}

	_, topicError := api.SetGroupTopic(publicChannel.ID, channelTopic)

	if topicError != nil {
		return topicError
	}

	d.SetId(publicChannel.ID)
	return nil
}
