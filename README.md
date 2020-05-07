# Terraform Provider for Slack [![CircleCI](https://circleci.com/gh/TimDurward/terraform-provider-slack/tree/master.svg?style=svg)](https://circleci.com/gh/TimDurward/terraform-provider-slack/tree/master)

This is a [Slack](https://slack.com) provider for [Terraform](https://www.terraform.io/)

_The provider allows for creation & desruction of public Slack channels_

# Installation

## Requirements

terraform-provider-slack is based on Terraform, this means that you need

* [Terraform](https://www.terraform.io/downloads.html) >=0.10.0

## Installation from binaries (recommended)

The recommended way to install terraform-provider-slack is use the binary distributions from the [Releases](https://github.com/TimDurward/terraform-provider-slack/releases) page. The packages are available for Linux and macOS.

Download the latest release for your OS.

### Linux

```sh
> wget https://github.com/TimDurward/terraform-provider-slack/releases/download/v0.2.0/terraform-provider-slack_linux_amd64
# Now copy the binary to the Terraform's plugins folder.
mkdir -p ~/.terraform.d/plugins/
mv terraform-provider-slack_linux_amd64 ~/.terraform.d/plugins/
```

### Mac

```sh
> wget https://github.com/TimDurward/terraform-provider-slack/releases/download/v0.2.0/terraform-provider-slack_darwin_amd64
# Now copy the binary to the Terraform's plugins folder.
mkdir -p ~/.terraform.d/plugins/
mv terraform-provider-slack_darwin_amd64 ~/.terraform.d/plugins/
```

## Compiling from source
Compile from source easily using [Makefile](https://github.com/TimDurward/terraform-provider-slack/blob/master/Makefile). 

\*_This is only necessary if your target OS/Architecture isn't listed in [releases](https://github.com/TimDurward/terraform-provider-slack/releases)_

```sh
make build

# Target is compiled at '$GOBIN/terraform-provider-slack'
mv $GOBIN/terraform-provider-slack ~/.terraform.d/plugins/
```

# Example

```hcl
provider "slack" {
  api_token = "SLACK_API_TOKEN"
}

resource "slack_channel" "jenkins_ci" {
  channel_name = "jenkins"
  channel_topic = "Jenkins Integration for production deploys"
}
```
