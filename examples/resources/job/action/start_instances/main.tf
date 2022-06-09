# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスを起動アクション
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のEC2インスタンスを指定
#     - タグのキー
#     - env
#   - タグの値
#     - production
# ----------------------------------------------------------
resource "cloudautomator_job" "example-start-instances-job" {
  name = "example-start-instances-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "start_instances"
  start_instances_action_value {
  region = "ap-northeast-1"
  specify_instance = "tag"
  tag_key = "env"
  tag_value = "production"
  trace_status = "true"
  status_checks_enable = "true"
  }
}
