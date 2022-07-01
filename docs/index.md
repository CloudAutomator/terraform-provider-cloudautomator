---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudautomator Provider"
subcategory: ""
description: |-

---

# cloudautomator Provider

## Authentication and Configuration
Configuration for the Cloud Automator Provider can be derived from several sources, which are applied in the following order:

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

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String) Cloud Automator API key. This can also be set via the CLOUD_AUTOMATOR_API_KEY environment variable.