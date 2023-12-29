# ----------------------------------------------------------
# - アクション
#   - IAM: ユーザーからポリシーをデタッチ
# - アクションの設定
#   - ユーザー名
#     - example-user
#   - IAMポリシーARN
#     - arn:aws:iam::123456789012:policy/example-policy
# ----------------------------------------------------------

resource "cloudautomator_job" "example-detach-user-policy-job" {
  name           = "example-detach-user-policy-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "detach_user_policy"
  detach_user_policy_action_value {
    user_name  = "example-user"
    policy_arn = "arn:aws:iam::123456789012:policy/example-policy"
  }
}
