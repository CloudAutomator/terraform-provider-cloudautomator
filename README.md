# Terraform Provider For Cloud Automator
Terraform provider for Cloud Automator

- Terraform Website: https://terraform.io
- Cloud Automator: https://cloudautomator.com
- Cloud Automator API Document: https://api.cloudautomator.com

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x

## Building The Provider

Clone repository to: $GOPATH/src/github.com/penta515/terraform-provider-cloudautomator

```sh
$ mkdir -p $GOPATH/src/github.com/penta515; cd $GOPATH/src/github.com/penta515
$ git clone git@github.com:penta515/terraform-provider-cloudautomator
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/penta515/terraform-provider-cloudautomator
$ make install
```

## Usage Example

```hcl
# Configure the Cloud Automator Provider
terraform {
  required_providers {
    cloudautomator = {
      source = "penta515/cloudautomator"

      version = "0.0.1"
    }
  }
}

provider "cloudautomator" {
  api_key = "abcdefghijklmnopqrstuvwxyz"
}

resource "cloudautomator_job" "example-job" {
    name = "example-job"

    group_id = 10
    aws_account_id = 20

    rule_type = "cron"
    cron_rule_value {
        hour = "9"
        minutes = "30"
        schedule_type = "weekly"
        weekly_schedule = [
            "monday",
            "sunday"
        ]
        time_zone = "Tokyo"
        dates_to_skip = ["2022-12-31"]
        national_holiday_schedule = "true"
    }

    action_type = "delay"
    delay_action_value {
        delay_minutes = 1
    }
}
```
