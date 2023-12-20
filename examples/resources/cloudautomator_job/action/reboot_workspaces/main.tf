# ----------------------------------------------------------
# - アクション
#   - ELB(CLB): EC2インスタンスを登録
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - インスタンス特定に利用するタグのキー
#     - env
#   - インスタンス特定に利用するタグの値
#     - production
#   - EC2インスタンスを登録するELB(CLB)名
#     - example-elb
# ----------------------------------------------------------

resource "cloudautomator_job" "example-register-instances-job" {
  name           = "example-register-instances-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "register_instances"
  register_instances_action_value {
    region             = "ap-northeast-1"
    specify_instance   = "tag"
    tag_key            = "env"
    tag_value          = "production"
    load_balancer_name = "example-elb"
  }
}
