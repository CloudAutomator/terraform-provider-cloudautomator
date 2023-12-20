# ----------------------------------------------------------
# - アクション
#   - ELB(CLB): EC2インスタンスを登録解除
# - アクションの設定
#   - 対象のELB(CLB)が存在するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のEC2インスタンスを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - EC2インスタンスを登録解除するELB(CLB)名
#     - test
# ----------------------------------------------------------

resource "cloudautomator_job" "example-deregister-instances-job" {
  name           = "example-deregister-instances-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "deregister_instances"
  deregister_instances_action_value {
    region             = "ap-northeast-1"
    specify_instance   = "tag"
    tag_key            = "env"
    tag_value          = "develop"
    load_balancer_name = "test"
  }
}
