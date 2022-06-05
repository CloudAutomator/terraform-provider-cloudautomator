variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - スケジュールトリガー
#   - 実行予定日時
#     - 2023-12-31 10:10:00
#     - 2023-01-01 22:40:00
# ----------------------------------------------------------
resource "cloudautomator_job" "schedule-start-instances" {
  name = "example-schedule-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "schedule"
  schedule_rule_value {
    schedule = "2023-12-31 10:10:00\n2023-01-01 22:40:00"
    time_zone = "Tokyo"
  }

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}
