# ----------------------------------------------------------
# - アクション
#   - WorkSpaces: WorkSpaceを再起動
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - WorkSpaceの特定に利用するタグのキー
#     - env
#   - WorkSpaceの特定に利用するタグの値
#     - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-reboot-workspaces-job" {
  name           = "example-reboot-workspaces-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "reboot_workspaces"
  reboot_workspaces_action_value {
    region    = "ap-northeast-1"
    tag_key   = "env"
    tag_value = "production"
  }
}
