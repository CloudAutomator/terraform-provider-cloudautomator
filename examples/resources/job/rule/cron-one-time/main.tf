# ----------------------------------------------------------
# - タイマートリガー
#   - 一度きり
#     - 2023-01-01
#     - 9:30
#   - ジョブの開始が遅延した場合に実行の開始をキャンセルする遅延時間
#     - キャンセルしない
# ----------------------------------------------------------
resource "cloudautomator_job" "cron-one-time-start-instances" {
  name = "example-cron-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "cron"
  cron_rule_value {
    hour = "9"
    minutes = "30"
    schedule_type = "one_time"
    one_time_schedule = "2023/01/01"
    time_zone = "Tokyo"
  }

  action_type = "delay"
  delay_action_value {
    delay_minutes = 1
  }
}
