# Terraform Slack Provider
# Example
provider "slack" {
  api_token = "SLACK_API_TOKEN"
}

resource "slack_channel" "jenkins_ci" {
  channel_name = "jenkins"
}
