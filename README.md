# Terraform Provider For Cloud Automator
Terraform provider for Cloud Automator

- Terraform Website: https://terraform.io
- Cloud Automator: https://cloudautomator.com
- Cloud Automator API Document: https://api.cloudautomator.com

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x

## Building The Provider

Clone repository to: $GOPATH/src/github.com/CloudAutomator/terraform-provider-cloudautomator

```sh
$ mkdir -p $GOPATH/src/github.com/CloudAutomator; cd $GOPATH/src/github.com/CloudAutomator
$ git clone git@github.com:CloudAutomator/terraform-provider-cloudautomator
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/CloudAutomator/terraform-provider-cloudautomator
$ make install VERSION=0.2.8
```

## Authentication and Configuration
Cloud Automator Provider authentication settings are applied in the following order.

1. Parameters in the provider configuration
1. Environment variables

### Provider Configuration

```hcl
provider "cloudautomator" {
  api_key = "abcdefghijklmnopqrstuvwxyz"
}
```

### Environment Variables

```hcl
provider "cloudautomator" {}
```

```shell
$ export CLOUD_AUTOMATOR_API_KEY="abcdefghijklmnopqrstuvwxyz"
$ terraform plan
```

## Custom Endpoint Configuration
Cloud Automator Provider can be customized to connect to non-default endpoints and is applied in the following order.

1. Parameters in the provider configuration
1. Environment variables

### Provider Configuration

```tf
provider "cloudautomator" {
  ...

  endpoint = "http://localhost:3000/api/v1"
}
```

### Environment Variables

```hcl
provider "cloudautomator" {}
```

```shell
$ export CLOUD_AUTOMATOR_API_ENDPOINT="http://localhost:3000/api/v1"
$ terraform plan
```

## Usage Example

```hcl
# Configure the Cloud Automator Provider
terraform {
  required_providers {
    cloudautomator = {
      source = "CloudAutomator/cloudautomator"
      version = "0.2.8"
    }
  }
}

provider "cloudautomator" {}

resource "cloudautomator_job" "example-job" {
  name = "example-job"

  group_id       = 10
  aws_account_id = 20

  rule_type = "cron"
  cron_rule_value {
    hour          = "9"
    minutes       = "30"
    schedule_type = "weekly"
    weekly_schedule = [
      "monday",
      "sunday"
    ]
    time_zone                 = "Tokyo"
    dates_to_skip             = ["2022-12-31"]
    national_holiday_schedule = "true"
  }

  action_type = "delay"
  delay_action_value {
    delay_minutes = 1
  }

  completed_post_process_id = [100]
  failed_post_process_id    = [200]
}
```
