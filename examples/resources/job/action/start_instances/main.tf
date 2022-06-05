variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

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
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

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
