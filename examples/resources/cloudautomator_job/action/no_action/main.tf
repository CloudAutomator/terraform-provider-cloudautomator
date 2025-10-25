# ----------------------------------------------------------
# - アクション: Other: アクションなし
# - アクションの設定:
#   - アクション値は不要
#   - ジョブワークフローのワークフロートリガージョブとして使用
# ----------------------------------------------------------

resource "cloudautomator_job" "example-no-action" {
  name     = "example-no-action"
  group_id = 10

  for_workflow = true

  rule_type = "webhook"

  action_type = "no_action"
}
