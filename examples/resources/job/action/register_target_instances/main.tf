# ----------------------------------------------------------
# - アクション
#   - ELB(ALB/NLB): ターゲットグループにEC2インスタンスを登録
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - 対象のターゲットグループのARN
#     - arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/example-elb/1234567890123456
#   - 登録するEC2インスタンスのタグのキー
#     - env
#   - 登録するEC2インスタンスのタグの値
#     - production
# ----------------------------------------------------------

resource "cloudautomator_job" "example-register-target-instances-job" {
  name           = "example-register-target-instances-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "register_target_instances"
  register_target_instances_action_value {
    region           = "ap-northeast-1"
    target_group_arn = "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/example-elb/1234567890123456"
    tag_key          = "env"
    tag_value        = "production"
  }
}
