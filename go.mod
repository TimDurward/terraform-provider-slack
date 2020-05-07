module github.com/TimDurward/terraform-provider-slack

go 1.12

require (
	github.com/hashicorp/terraform v0.12.20
	github.com/slack-go/slack v0.6.4
)

replace github.com/slack-go/slack => github.com/timdurward/slack v0.6.5
