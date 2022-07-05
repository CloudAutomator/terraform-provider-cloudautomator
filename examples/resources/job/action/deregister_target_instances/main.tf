# ----------------------------------------------------------
# - アクション
#   - ELB(ALB/NLB): ターゲットグループからEC2インスタンスを登録解除
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - 対象のターゲットグループのARN
#     - arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:loadbalancer/app/example/1234567890123456
#   - 登録解除するEC2インスタンスのタグのキー
#     - env
#   - 登録解除するEC2インスタンスのタグの値
#     - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-deregister-target-instances-job" {
  name           = "example-deregister-target-instances-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "deregister_target_instances"
  deregister_target_instances_action_value {
    region           = "ap-northeast-1"
    target_group_arn = "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:loadbalancer/app/example/1234567890123456"
    tag_key          = "env"
    tag_value        = "production"
  }
}
