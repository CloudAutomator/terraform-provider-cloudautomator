variable "group_id" {
  default = 123
}

data "cloudautomator_aws_account" "production" {
  group_id = var.group_id
  id = 456
}

# ----------------------------------------------------------
# - アクション
#   - EC2: AMIをリージョン間でコピー
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のEC2インスタンスを指定
#     - タグのキー
#       - env
#     - タグの値
#       - production
#   - AMIのコピー完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------
resource "cloudautomator_job" "example-copy-image-job" {
	name = "example-copy-image-job"
  group_id = var.group_id
  aws_account_id = data.cloudautomator_aws_account.production.id

	rule_type = "webhook"

	action_type = "copy_image"
	copy_image_action_value {
		source_region = "ap-northeast-1"
		destination_region = "us-east-1"
		specify_image = "tag"
		tag_key = "env"
		tag_value = "develop"
		trace_status = "true"
	}
}
