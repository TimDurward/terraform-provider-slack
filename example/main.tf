# Custom Terraform Provider: Slack
# Required: API_TOKEN

provider "slack" {
  api_token = "SLACK_API_TOKEN"
}

# Create Slack Channel named OG Channel
resource "slack_channel" "this" {
  channel_name = "OG Channel"
}