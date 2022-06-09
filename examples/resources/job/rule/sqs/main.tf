# ----------------------------------------------------------
# - SQSトリガー
#   - SQSのキューを検索する際に利用するAWSアカウントのID
#     - 1
#   - SQSのキューが存在するリージョン名
#     - ap-northeast-1
#   - SQSのキュー名
#     - test-queue
# ----------------------------------------------------------
resource "cloudautomator_job" "sqs-schedule-start-instances" {
  name = "example-sqs-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "sqs_v2"
  sqs_v2_rule_value {
    sqs_aws_account_id = 20
    sqs_region = "ap-northeast-1"
    queue = "test-queue"
  }

  action_type = "delay"
  delay_action_value {
    delay_minutes = 1
  }
}
