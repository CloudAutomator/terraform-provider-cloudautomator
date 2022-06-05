variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスタイプを変更
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のDBインスタンスを指定
#     - タグのキー
#     - env
#   - タグの値
#     - production
#   - DBインスタンスクラス
#     - db.t3.micro
# ----------------------------------------------------------
resource "cloudautomator_job" "example-change-rds-instance-class-job" {
  name = "example-change-rds-instance-class-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

  rule_type = "immediate_execution"

	action_type = "change_rds_instance_class"
	change_rds_instance_class_action_value {
		region = "ap-northeast-1"
		specify_rds_instance = "tag"
		tag_key = "env"
		tag_value = "develop"
		db_instance_class = "db.t3.micro"
	}
}
