# ----------------------------------------------------------
# - アクション
#   - WorkSpaces: WorkSpaceを削除
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のWorkSpaceを指定
#     - WorkSpace特定に利用するタグのキー
#       - env
#     - WorkSpace特定に利用するタグの値
#       - production
#   - WorkSpaceの作成完了をジョブ完了の判定にするフラグ
#     - する
# ----------------------------------------------------------

resource "cloudautomator_job" "example-terminate-workspaces-job" {
  name           = "example-terminate-workspaces-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "terminate_workspaces"
  terminate_workspaces_action_value {
    region            = "ap-northeast-1"
    specify_workspace = "tag"
    tag_key           = "env"
    tag_value         = "production"
    trace_status      = "true"
  }
}
