# Terraform Provider for Slack [![CircleCI](https://circleci.com/gh/TimDurward/terraform-provider-slack/tree/master.svg?style=svg)](https://circleci.com/gh/TimDurward/terraform-provider-slack/tree/master)

This is a [Slack](https://slack.com) provider for [Terraform](https://www.terraform.io/)

_The provider allows for creation & desruction of public Slack channels_

# Installation

## Requirements

terraform-provider-slack is based on Terraform, this means that you need

* [Terraform](https://www.terraform.io/downloads.html) >=0.12.0

## Installation from binaries (recommended)

The recommended way to install terraform-provider-slack is use the binary distributions from the [Releases](https://github.com/TimDurward/terraform-provider-slack/releases) page. The packages are available for Linux and macOS.

Download and uncompress the latest release for your OS. This example uses the linux binary.

```sh
> wget https://github.com/TimDurward/terraform-provider-slack/releases/download/v0.1.0/terraform-provider-slack_v0.1.0_darwin_amd64.tar.gz
> tar -xvf terraform-provider-slack*.tar.gz
```

Now copy the binary to the Terraform's plugins folder, if is your first plugin maybe isn't present.

```sh
> mkdir -p ~/.terraform.d/plugins/
> mv terraform-provider-slack*/terraform-provider-slack ~/.terraform.d/plugins/
```

### \*_You can always install from source with [Makefile](https://github.com/TimDurward/terraform-provider-slack/blob/master/Makefile)_

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
