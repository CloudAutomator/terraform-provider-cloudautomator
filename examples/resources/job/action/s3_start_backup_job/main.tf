# ----------------------------------------------------------
# - アクション
#   - S3: バケットをバックアップ
# - アクションの設定
#   - 対象のバケットとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のバケットの名前
#     - test-bucket
#   - バックアップボールトの名前
#     - TestBackup
#   - ライフサイクル-
#     - 7日
#   - バックアップ取得時に使うIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForBackup
#   - 作成した復旧ポイントに割り当てるタグの配列
#     - key1: value1
#     - key2: value2
# ----------------------------------------------------------

resource "cloudautomator_job" "example-s3-start-backup-job-job" {
  name           = "example-s3-start-backup-job-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "s3_start_backup_job"
  s3_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    bucket_name                 = "test-bucket"
    backup_vault_name           = "TestBackup"
    lifecycle_delete_after_days = 7
    iam_role_arn                = "arn:aws:iam::123456789012:role/RoleForBackup"
    additional_tags {
      key   = "key1"
      value = "value1"
    }
    additional_tags {
      key   = "key2"
      value = "value2"
    }
  }
}
