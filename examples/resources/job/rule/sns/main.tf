variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - SNSトリガー
#   - 実行予定日時
#     - 2023-12-31 10:10:00
#     - 2023-01-01 22:40:00
# ----------------------------------------------------------
resource "cloudautomator_job" "sns-start-instances" {
  name = "example-sns-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "amazon_sns"

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}
