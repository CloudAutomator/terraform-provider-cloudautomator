# ----------------------------------------------------------
# - アクション
#   - RDS: DBインスタンスを再起動
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のDBインスタンスを指定
#     - タグのキー
#       - env
#     - タグの値
#       - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-reboot-rds-instances-job" {
  name           = "example-reboot-rds-instances-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "reboot_rds_instances"
  reboot_rds_instances_action_value {
    region               = "ap-northeast-1"
    specify_rds_instance = "tag"
    tag_key              = "env"
    tag_value            = "production"
  }
}
