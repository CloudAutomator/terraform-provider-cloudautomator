variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - タイマートリガー
#   - 毎月(曜日指定)
#     - 最終金曜日
#     - 9:00
#   - 実行日と日本の祝日が重なっていた場合
#     - スキップしない
#   - ジョブ実行をスキップする日付
#     - 2023-01-01
#     - 2023-01-02
#     - 2022-01-03
#   - ジョブの開始が遅延した場合に実行の開始をキャンセルする遅延時間
#     - 30分
# ----------------------------------------------------------
resource "cloudautomator_job" "cron-monthly-day-of-week-schedule-start-instances" {
  name = "example-cron-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "cron"
  cron_rule_value {
    hour = "9"
    minutes = "30"
    schedule_type = "monthly_day_of_week"
    monthly_day_of_week_schedule {
      friday = [-1]
    }
    national_holiday_schedule = "false"
    dates_to_skip = [
      "2023-01-01",
      "2023-01-02",
      "2023-01-03"
    ]
    start_timeout_minutes = "30"
    time_zone = "Tokyo"
  }

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}
