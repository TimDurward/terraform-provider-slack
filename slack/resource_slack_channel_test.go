package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccResourceSlackChannel_Create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t)},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestSlackChannelConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("resource.slack_channel.jenkins_ci", "channel_name"),
					resource.TestCheckResourceAttrSet("resource.slack_channel.jenkins_ci", "channel_topic"),
					resource.TestCheckResourceAttr("resource.slack_channel.jenkins_ci","channel_name", "jenkins"),
					resource.TestCheckResourceAttr("resource.slack_channel.jenkins_ci","channel_name", "Jenkins Integration for production deploys"),
				),
			},
		},
	})
}

const TestSlackChannelConfig = `
resource "slack_channel" "jenkins_ci" {
  channel_name = "jenkins"
  channel_topic = "Jenkins Integration for production deploys"
}
`