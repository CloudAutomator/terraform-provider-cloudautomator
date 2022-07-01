# ----------------------------------------------------------
# - アクション
#   - Other: 指定時間待機
# - アクションの設定
#   - 待機する時間
#     - 30分
# ----------------------------------------------------------

resource "cloudautomator_job" "example-delay-job" {
  name           = "example-delay-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "delay"
  delay_action_value {
    delay_minutes = 30
  }
}
