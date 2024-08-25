# ----------------------------------------------------------
# - アクション
#   - EC2: 過去のAMIとスナップショットをまとめて削除
# - アクションの設定
#   - 特定のタグが付いたAMIを除外するかどうか
#     - 除外する
#   - 除外するAMIの特定に利用するタグのキー
#     - env
#   - 除外するAMIの特定に利用するタグの値
#     - production
#   - 削除するAMIの指定方法
#     - 日数で指定する
#   - 削除するAMIを日数で指定する場合の日数
#     - 365
# ----------------------------------------------------------

resource "cloudautomator_job" "example-bulk-delete-images-job" {
  name            = "example-bulk-delete-images-job"
  group_id        = 10
  aws_account_ids = [20]

  rule_type = "immediate_execution"

  action_type = "bulk_delete_images"
  bulk_delete_images_action_value {
    exclude_by_tag_bulk_delete_images       = true
    exclude_by_tag_key_bulk_delete_images   = "env"
    exclude_by_tag_value_bulk_delete_images = "production"
    specify_base_date                       = "before_days"
    before_days                             = 365
  }
}
