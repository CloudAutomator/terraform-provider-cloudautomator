# ----------------------------------------------------------
# - アクション
#   - EC2: EBSスナップショットをリージョン間でコピー
# - アクションの設定
#   - EBSスナップショットのコピー元のAWSリージョン
#     - ap-northeast-1
#   - EBSスナップショットのコピー先のAWSリージョン
#     - us-east-1
#   - タグで対象のEBSスナップショットを指定
#     - タグのキー
#       - env
#     - タグの値
#       - production
#   - DBインスタンスクラス
#     - db.t3.micro
#   - EBSスナップショットのコピー完了をジョブ完了の判定にする
# ----------------------------------------------------------
resource "cloudautomator_job" "example-change-rds-instance-class-job" {
  name = "example-change-rds-instance-class-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "copy_ebs_snapshot"
  copy_ebs_snapshot_action_value {
    source_region = "ap-northeast-1"
    destination_region = "us-east-1"
    specify_ebs_snapshot = "tag"
    tag_key = "env"
    tag_value = "develop"
    trace_status = "true"
  }
}
