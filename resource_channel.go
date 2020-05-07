package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/slack-go/slack"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Exists: resourceChannelExists,
		Importer: &schema.ResourceImporter{
			State: resourceChannelImportState,
		},

		Schema: map[string]*schema.Schema{
			"channel_name": &schema.Schema{
				Type:         schema.TypeString,
				Description:  "The name of Slack Channel that will be created",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 80),
			},

			"channel_purpose": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Sets the purpose for a channel",
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
	api := slack.New(meta.(*Config).APIToken)

	// Create Slack Channel
	channel, err := api.CreateChannel(d.Get("channel_name").(string))
	if err != nil {
		return err
	}
	d.SetId(channel.ID)

	// Update Slack Channel Purpose
	if _, err := api.SetChannelPurpose(d.Id(), d.Get("channel_purpose").(string)); err != nil {
		return err
	}
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
	api := slack.New(meta.(*Config).APIToken)

	name := d.Get("channel_name").(string)
	if _, err := api.RenameChannel(d.Id(), name); err != nil {
		return err
	}
	// Update Slack Channel Purpose
	if _, err := api.SetChannelPurpose(d.Id(), d.Get("channel_purpose").(string)); err != nil {
		return err
	}
	// Update Slack Channel Topic
	if _, err := api.SetChannelTopic(d.Id(), d.Get("channel_topic").(string)); err != nil {
		return err
	}
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	// Deletes Slack Channel and clears state
	if _, err := api.DeleteChannel(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceChannelImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	api := slack.New(meta.(*Config).APIToken)

	// Checks if Slack Channel exists, if not remove resource from state
	channel, err := api.GetConversationInfo(d.Id(), false)
	if err != nil {
		d.SetId("")
		return nil, err
	}
	d.Set("channel_name", channel.Name)
	d.Set("channel_purpose", channel.Purpose.Value)
	d.Set("channel_topic", channel.Topic.Value)

	return []*schema.ResourceData{d}, nil
}
