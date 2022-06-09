# ----------------------------------------------------------
# - 手動トリガー
# ----------------------------------------------------------
resource "cloudautomator_job" "immediate-execution-start-instances" {
  name = "example-immediate-execution-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "delay"
  delay_action_value {
    delay_minutes = 1
  }
}
