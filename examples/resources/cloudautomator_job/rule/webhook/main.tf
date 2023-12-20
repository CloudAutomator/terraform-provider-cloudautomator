# ----------------------------------------------------------
# - HTTPトリガー
# ----------------------------------------------------------

resource "cloudautomator_job" "webhook-start-instances" {
  name           = "example-webhook-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "delay"
  delay_action_value {
    delay_minutes = 1
  }
}
