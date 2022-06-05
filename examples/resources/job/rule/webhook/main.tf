variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - HTTPトリガー
# ----------------------------------------------------------
resource "cloudautomator_job" "webhook-start-instances" {
  name = "example-webhook-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "webhook"

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}
