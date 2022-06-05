variable "group_id" {
  default = 123
}

# ----------------------------------------------------------
# - Slack後処理
#   - グループID
#     - 123
#   - 通知するチャンネル名
#     - test
#   - 通知内容の言語
#     - ja
#   - 通知内容のタイムゾーン
#     - Tokyo
# ----------------------------------------------------------
resource "cloudautomator_post_process" "slack" {
  name = "slack"
  group_id = var.group_id
  service = "slack"

  slack_parameters {
    slack_channel_name = "test"
    slack_language = "ja"
    slack_time_zone = "Tokyo"
  }
}
