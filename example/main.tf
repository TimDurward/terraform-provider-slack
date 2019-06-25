# Terraform Slack Provider
# Example
provider "slack" {
  api_token = "SLACK_API_TOKEN"
}

resource "slack_channel" "jenkins_ci" {
  channel_name = "jenkins"
}

data "slack_user" "jenkins_data"{
  name = "user"
  email = "user@mail.com"
}

output "channel_id" {
  value = slack_channel.jenkins_ci.id
}
