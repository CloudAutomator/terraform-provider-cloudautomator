# ----------------------------------------------------------
# - タイマートリガー
#   - 毎週
#     - 月曜日、日曜日
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
resource "cloudautomator_job" "cron-weekly-start-instances" {
  name = "example-cron-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "cron"
  cron_rule_value {
    hour = "9"
    minutes = "00"
    schedule_type = "weekly"
    weekly_schedule = [
      "monday",
      "sunday"
    ]
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
