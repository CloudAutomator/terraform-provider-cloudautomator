# ----------------------------------------------------------
# - アクション
#   - EFS: ファイルシステムをバックアップ
# - アクションの設定
#   - 対象のEFSファイルシステムとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のEFSファイルシステムのID
#     - fs-abcdefg1234567890
#   - 対象のバックアップボールトの名前
#     - example-backup
#   - バックアップの保持期間(日数)
#     - 無期限
#   - バックアップ取得時に使うIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForBackup
# ----------------------------------------------------------

resource "cloudautomator_job" "example-efs-start-backup-job" {
  name           = "example-efs-start-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "efs_start_backup_job"
  efs_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    file_system_id              = "fs-abcdefg1234567890"
    backup_vault_name           = "example-backup"
    lifecycle_delete_after_days = null
    iam_role_arn                = "arn:aws:iam::123456789012:role/RoleForBackup"
  }
}

# ----------------------------------------------------------
# - アクション
#   - EFS: ファイルシステムをバックアップ
# - アクションの設定
#   - 対象のインスタンスとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のバックアップボールトの名前
#     - example-backup
#   - バックアップの保持期間(日数)
#     - 7
#   - バックアップ取得時に使うIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForBackup
#   - 作成した復旧ポイントに割り当てるタグの配列
#     - key1: value1
#     - key2: value2
#     - key3: value3
# ----------------------------------------------------------

resource "cloudautomator_job" "example-efs-start-backup-job" {
  name           = "example-efs-start-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "efs_start_backup_job"
  efs_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    file_system_id              = "fs-abcdefg1234567890"
    backup_vault_name           = "example-backup"
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

    additional_tags {
      key   = "key3"
      value = "value3"
    }
  }
}
