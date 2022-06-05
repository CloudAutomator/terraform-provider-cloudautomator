variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

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
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "sqs_v2"
  sqs_v2_rule_value {
    sqs_aws_account_id = data.cloudautomator_aws_account.production.id
    sqs_region = "ap-northeast-1"
    queue = "test-queue"
  }

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}
