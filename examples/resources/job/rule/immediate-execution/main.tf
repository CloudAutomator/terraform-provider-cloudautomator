variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - 手動トリガー
# ----------------------------------------------------------
resource "cloudautomator_job" "immediate-execution-start-instances" {
  name = "example-immediate-execution-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "immediate_execution"

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}
