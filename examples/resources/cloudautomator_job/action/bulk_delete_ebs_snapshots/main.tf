# ----------------------------------------------------------
# - アクション
#   - EC2: 過去のEBSスナップショットをまとめて削除
# - アクションの設定
#   - 特定のタグが付いたEBSスナップショットを除外するかどうか
#     - 除外する
#   - 除外するEBSスナップショットの特定に利用するタグのキー
#     - env
#   - 除外するEBSスナップショットの特定に利用するタグの値
#     - production
#   - 削除するEBSスナップショットの指定方法
#     - 日数で指定する
#   - 削除するEBSスナップショットを日数で指定する場合の日数
#     - 365
# ----------------------------------------------------------

resource "cloudautomator_job" "example-bulk_delete_ebs_snapshots-job" {
  name            = "example-bulk_delete_ebs_snapshots-job"
  group_id        = 10
  aws_account_ids = [20]

  rule_type = "immediate_execution"

  action_type = "bulk_delete_ebs_snapshots"
  bulk_delete_ebs_snapshots_action_value {
    exclude_by_tag_bulk_delete_ebs_snapshots       = true
    exclude_by_tag_key_bulk_delete_ebs_snapshots   = "env"
    exclude_by_tag_value_bulk_delete_ebs_snapshots = "production"
    specify_base_date                              = "before_days"
    before_days                                    = 365
  }
}
