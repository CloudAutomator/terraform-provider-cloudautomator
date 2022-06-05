variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - アクション
#   - RDS: DBスナップショットをリージョン間でコピー
# - アクションの設定
#   - DBスナップショットのコピー元のAWSリージョン
#     - ap-northeast-1
#   - DBスナップショットのコピー先のAWSリージョン
#     - us-east-1
#   - IDで対象のDBスナップショットを指定
#     - 対象のDBスナップショットID
#       - test-db
#   - コピー先リージョンに設定するオプショングループ名
#     - default:mysql-5-6
#   - DBスナップショットのコピー完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------
resource "cloudautomator_job" "example-copy-rds-snapshot-job" {
	name = "example-copy-rds-snapshot-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

	rule_type = "webhook"

	action_type = "copy_rds_snapshot"
	copy_rds_snapshot_action_value {
		source_region = "ap-northeast-1"
		destination_region = "us-east-1"
		specify_rds_snapshot = "identifier"
		rds_snapshot_id = "test-db"
		option_group_name = "default:mysql-5-6"
		trace_status = "true"
	}
}
