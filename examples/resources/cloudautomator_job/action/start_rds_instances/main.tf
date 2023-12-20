# ----------------------------------------------------------
# - アクション
#   - RDS: DBインスタンスを起動
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のRDSインスタンスを指定
#     - RDSインスタンス特定に利用するタグのキー
#       - env
#     - RDSインスタンス特定に利用するタグの値
#       - production
#   - RDSインスタンスの起動完了をジョブ完了の判定にするフラグ
#     - true
# ----------------------------------------------------------

resource "cloudautomator_job" "example-start-rds-instances-job" {
  name           = "example-start-rds-instances-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "start_rds_instances"
  start_rds_instances_action_value {
    region               = "ap-northeast-1"
    specify_rds_instance = "tag"
    tag_key              = "env"
    tag_value            = "production"
    trace_status         = "true"
  }
}
