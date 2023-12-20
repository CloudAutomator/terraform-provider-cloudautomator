# ----------------------------------------------------------
# - SQS後処理
#   - グループID
#     - 123
#   - SQSのキューを検索する際に利用するAWSアカウントのID
#     - 1
#   - SQSのキュー名
#     - test-queue
#   - SQSのキューが存在するリージョン名
#     - ap-northeast-1
# ----------------------------------------------------------

resource "cloudautomator_post_process" "sqs" {
  name     = "sqs"
  group_id = 10
  service  = "sqs"

  sqs_parameters {
    sqs_aws_account_id = var.group_id
    sqs_queue          = "test-queue"
    sqs_region         = "ap-northeast-1"
  }
}
