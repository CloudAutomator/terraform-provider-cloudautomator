# ----------------------------------------------------------
# - アクション
#   - WorkSpaces: WorkSpaceを起動
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - WorkSpace特定に利用するタグのキー
#     - env
#   - WorkSpace特定に利用するタグの値
#     - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-start-workspaces-job" {
  name           = "example-start-workspaces-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "start_workspaces"
  start_workspaces_action_value {
    region    = "ap-northeast-1"
    tag_key   = "env"
    tag_value = "production"
  }
}
