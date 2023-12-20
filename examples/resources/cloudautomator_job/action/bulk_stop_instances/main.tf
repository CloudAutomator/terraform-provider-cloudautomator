# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスをすべて停止
# - アクションの設定
#   - 特定のタグが付いたインスタンスを除外するかどうか
#     - 除外する
#   - 除外するインスタンスの特定に利用するタグのキー
#     - env
#   - 除外するインスタンスの特定に利用するタグの値
#     - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-bulk-stop-instances-job" {
  name            = "example-bulk-stop-instances-job"
  group_id        = 10
  aws_account_ids = [20]

  rule_type = "immediate_execution"

  action_type = "bulk_stop_instances"
  bulk_stop_instances_action_value {
    exclude_by_tag       = true
    exclude_by_tag_key   = "env"
    exclude_by_tag_value = "develop"
  }
}
