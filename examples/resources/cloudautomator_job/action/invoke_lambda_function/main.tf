# ----------------------------------------------------------
# - アクション
#   - Lambda: 関数を呼び出し
# - アクションの設定
#   - AWSリージョン
#     - ap-northeast-1
#   - Lambda関数名
#     - test-function
#   - イベントJSON
#     - {}
# ----------------------------------------------------------

resource "cloudautomator_job" "example-invoke-lambda-function-job" {
  name           = "example-invoke-lambda-function-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "invoke_lambda_function"
  invoke_lambda_function_action_value {
    region        = "ap-northeast-1"
    function_name = "test-function"
    payload       = "{}"
  }
}
