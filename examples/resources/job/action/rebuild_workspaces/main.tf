# ----------------------------------------------------------
# - アクション
#   - WorkSpaces: WorkSpaceを再構築
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグのキー
#     - env
#   - タグの値
#     - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-rebuild-workspaces-job" {
  name           = "example-rebuild-workspaces-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "rebuild_workspaces"
  rebuild_workspaces_action_value {
    region    = "ap-northeast-1"
    tag_key   = "env"
    tag_value = "production"
  }
}
