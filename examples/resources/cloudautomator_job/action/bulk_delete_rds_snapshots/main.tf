# ----------------------------------------------------------
# - アクション
#   - RDS: 過去のDBスナップショットをまとめて削除
# - アクションの設定
#   - 特定のタグが付いたDBスナップショットを除外するかどうか
#     - 除外する
#   - 除外するDBスナップショットの特定に利用するタグのキー
#     - env
#   - 除外するDBスナップショットの特定に利用するタグの値
#     - production
#   - 削除するDBスナップショットの指定方法
#     - 日数で指定する
#   - 削除するDBスナップショットを日数で指定する場合の日数
#     - 365
# ----------------------------------------------------------

resource "cloudautomator_job" "example-bulk-delete-rds-snapshots-job" {
  name            = "example-bulk-delete-rds-snapshots-job"
  group_id        = 10
  aws_account_ids = [20]

  rule_type = "immediate_execution"

  action_type = "bulk_delete_rds_snapshots"
  bulk_delete_rds_snapshots_action_value {
    exclude_by_tag_bulk_delete_rds_snapshots       = true
    exclude_by_tag_key_bulk_delete_rds_snapshots   = "env"
    exclude_by_tag_value_bulk_delete_rds_snapshots = "production"
    specify_base_date                              = "before_days"
    before_days                                    = 365
  }
}
