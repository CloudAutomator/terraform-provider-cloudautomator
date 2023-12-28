# ----------------------------------------------------------
# - アクション
#   - IAM: ユーザーにポリシーをアタッチ
# - アクションの設定
#   - ユーザー名
#     - example-user
#   - IAMポリシーARN
#     - arn:aws:iam::123456789012:policy/example-policy
# ----------------------------------------------------------

resource "cloudautomator_job" "example-attach-user-policy-job" {
  name           = "example-attach-user-policy-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "attach_user_policy"
  attach_user_policy_action_value {
    user_name  = "example-user"
    policy_arn = "arn:aws:iam::123456789012:policy/example-policy"
  }
}
