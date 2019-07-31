package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
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
				Type:         schema.TypeString,
				Description:  "The name of Slack Channel that will be created",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 21),
			},

			"channel_topic": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Sets the topic for a channel",
				Optional:    true,
			},
			"force_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Default:     true,
				Description: "Force the deletion of the channel instead of archiving it",
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
	if err.Error() == "name_taken" {
		// Channel most likely has to be unarchived
		channels, err := api.GetChannels(false)
		if err != nil {
			return err
		}
		for _, c := range channels {
			if !c.IsArchived {
				continue
			}
			if c.Name != d.Get("channel_name").(string) {
				continue
			}
			channel, err := api.GetChannelInfo(c.ID)
			if err != nil {
				return err
			}
			err = api.UnarchiveChannel(channel.ID)
			d.SetId(channel.ID)
			if err != nil {
				return err
			}
		}
	} else if err == nil {
		d.SetId(channel.ID)
	} else if err != nil {
		return err
	}

	// Create Slack Channel Topic
	if _, err := api.SetChannelTopic(d.Id(), d.Get("channel_topic").(string)); err != nil {
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
	return nil
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIToken)

	if d.Get("force_delete").(bool) {
		// Deletes Slack Channel and clears state
		if _, err := api.DeleteChannel(d.Id()); err != nil {
			return err
		}
	} else {
		// Archives Slack Channel
		if err := api.ArchiveChannel(d.Id()); err != nil {
			return err
		}
	}

	return nil
}
