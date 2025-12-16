# ----------------------------------------------------------
# - アクション
#   - Backup: ボールトの復旧ポイントをコピー
# - アクションの設定
#   - コピー元のバックアップボールトが配置されているリージョン
#     - ap-northeast-1
#   - コピー元バックアップボールト名
#     - source-backup-vault
#   - 復旧ポイントをコピーするリソースの種類
#     - EC2
#   - 復旧ポイントをコピーするリソースのID
#     - i-00000001
#   - コピーを作成するときに使用するIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForCopy
#   - コピー先AWSアカウントの指定方法
#     - same (同一AWSアカウントにコピー)
#   - コピー先バックアップボールトがあるリージョン
#     - ap-southeast-1
#   - コピー先バックアップボールト名
#     - dest-backup-vault
#   - コピー保存期間（日数）
#     - 7
# ----------------------------------------------------------

resource "cloudautomator_job" "example-vault-recovery-point-start-copy-job-same-account" {
  name           = "example-vault-recovery-point-start-copy-job-same-account"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "vault_recovery_point_start_copy_job"
  vault_recovery_point_start_copy_job_action_value {
    source_region                   = "ap-northeast-1"
    source_backup_vault_name        = "source-backup-vault"
    resource_type                   = "EC2"
    resource_id                     = "i-00000001"
    iam_role_arn                    = "arn:aws:iam::123456789012:role/RoleForCopy"
    specify_destination_aws_account = "same"
    destination_region              = "ap-southeast-1"
    destination_backup_vault_name   = "dest-backup-vault"
    lifecycle_delete_after_days     = 7
  }
}

# ----------------------------------------------------------
# - アクション
#   - Backup: ボールトの復旧ポイントをコピー
# - アクションの設定
#   - コピー元のバックアップボールトが配置されているリージョン
#     - ap-northeast-1
#   - コピー元バックアップボールト名
#     - source-backup-vault
#   - 復旧ポイントをコピーするリソースの種類
#     - RDS (Aurora)
#   - 復旧ポイントをコピーするリソースのID
#     - production-database-1
#   - コピーを作成するときに使用するIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForCopy
#   - コピー先AWSアカウントの指定方法
#     - different (異なるAWSアカウントにコピー)
#   - コピー先バックアップボールトのARN
#     - arn:aws:backup:ap-northeast-1:987654321098:backup-vault:dest-backup-vault
#   - コピー保存期間（日数）
#     - 無期限
# ----------------------------------------------------------

resource "cloudautomator_job" "example-vault-recovery-point-start-copy-job-different-account" {
  name           = "example-vault-recovery-point-start-copy-job-different-account"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "vault_recovery_point_start_copy_job"
  vault_recovery_point_start_copy_job_action_value {
    source_region                   = "ap-northeast-1"
    source_backup_vault_name        = "source-backup-vault"
    resource_type                   = "RDS (Aurora)"
    resource_id                     = "production-database-1"
    iam_role_arn                    = "arn:aws:iam::123456789012:role/RoleForCopy"
    specify_destination_aws_account = "different"
    destination_backup_vault_arn    = "arn:aws:backup:ap-northeast-1:987654321098:backup-vault:dest-backup-vault"
    lifecycle_delete_after_days     = null
  }
}
