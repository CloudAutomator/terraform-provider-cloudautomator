# Terraform Provider For Cloud Automator
Terraform provider for Cloud Automator

- Terraform Website: https://terraform.io
- Cloud Automator: https://cloudautomator.com
- Cloud Automator API Document: https://api.cloudautomator.com

## Building The Provider

### Prerequisites

- Go 1.25 or later
- GNU Make

Clone the repository:

```sh
$ git clone git@github.com:CloudAutomator/terraform-provider-cloudautomator.git
$ cd terraform-provider-cloudautomator
```

Build the provider binary. The build writes `terraform-provider-cloudautomator_v<VERSION>` into `./bin/`:

```sh
$ make build VERSION=<release-version>
```

Install the built binary into the Terraform plugin directory for your OS/architecture:

```sh
$ make install VERSION=<release-version>
```

## Development Commands

The project uses a `GNUmakefile` for common development tasks.

### Testing

Run unit tests:

```sh
$ make test
```

Run acceptance tests (requires valid Cloud Automator API credentials):

```sh
$ make testacc
```

### Code Formatting

Format Terraform example files:

```sh
$ make fmt
```

### Documentation

Generate provider documentation from schema definitions:

```sh
$ make docs-generate
```

Verify documentation is up to date:

```sh
$ make test-docs
```

### Cleanup

Remove installed provider binaries:

```sh
$ make clean VERSION=<release-version>
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

  api_endpoint = "http://localhost:3000/api/v1/"
}
```

### Environment Variables

```hcl
provider "cloudautomator" {}
```

```shell
$ export CLOUD_AUTOMATOR_API_ENDPOINT="http://localhost:3000/api/v1/"
$ terraform plan
```

## Usage Example

```hcl
# Configure the Cloud Automator Provider
terraform {
  required_providers {
    cloudautomator = {
      source = "CloudAutomator/cloudautomator"
      version = "0.3.1"
    }
  }
}

provider "cloudautomator" {}

resource "cloudautomator_job" "example-job" {
  name = "example-job"

  group_id       = 10

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

## Provider Documentation

The provider ships with auto-generated documentation under `docs/`. Run `make docs-generate` after updating schemas to refresh the content.

For detailed usage examples of specific actions, see the `examples/` directory which contains over 70 sample configurations for various Cloud Automator job actions.

### Resources

- `cloudautomator_job` – Manage Cloud Automator jobs. See [`docs/resources/job.md`](docs/resources/job.md) for the complete schema.
- `cloudautomator_job_workflow` – Manage job workflows. See [`docs/resources/job_workflow.md`](docs/resources/job_workflow.md).
- `cloudautomator_post_process` – Manage post-process definitions. See [`docs/resources/post_process.md`](docs/resources/post_process.md).

### Data Sources

- `cloudautomator_job` – Retrieve job details. See [`docs/data-sources/job.md`](docs/data-sources/job.md).
- `cloudautomator_job_workflow` – Retrieve job workflow details. See [`docs/data-sources/job_workflow.md`](docs/data-sources/job_workflow.md).
- `cloudautomator_post_process` – Retrieve post-process definitions. See [`docs/data-sources/post_process.md`](docs/data-sources/post_process.md).
- `cloudautomator_aws_account` – Retrieve Cloud Automator AWS account metadata. See [`docs/data-sources/aws_account.md`](docs/data-sources/aws_account.md).
