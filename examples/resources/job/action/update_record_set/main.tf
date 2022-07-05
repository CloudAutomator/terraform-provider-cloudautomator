# ----------------------------------------------------------
# - アクション
#   - Route 53: リソースレコードセットを更新
# - アクションの設定
#   - リソースレコードセットを更新するホストゾーン
#     - test.local.
#   - 更新対象のリソースレコードセット
#     - aaa.test.local.
#   - リソースレコードタイプ
#     - A
#   - リソースレコードセットの値
#     - 1.2.3.4
# ----------------------------------------------------------

resource "cloudautomator_job" "example-update-record-set-job" {
  name           = "example-update-record-set-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "update_record_set"
  update_record_set_action_value {
    zone_name        = "test.local."
    record_set_name  = "aaa.test.local."
    record_set_type  = "A"
    record_set_value = "1.2.3.4"
  }
}
