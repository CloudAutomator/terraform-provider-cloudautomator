# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスをバックアップ
# - アクションの設定
#   - 対象のインスタンスとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のインスタンスのID
#     - i-00000001
#   - 対象のバックアップボールトの名前
#     - example-backup
#   - バックアップの保持期間(日数)
#     - 無期限
#   - バックアップ取得時に使うIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForBackup
# ----------------------------------------------------------

resource "cloudautomator_job" "example-ec2-start-backup-job" {
  name           = "example-ec2-start-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "ec2_start_backup_job"
  ec2_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    specify_instance            = "identifier"
    instance_id                 = "i-00000001"
    backup_vault_name           = "example-backup"
    lifecycle_delete_after_days = null
    iam_role_arn                = "arn:aws:iam::123456789012:role/RoleForBackup"
  }
}

# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスをバックアップ
# - アクションの設定
#   - 対象のインスタンスとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のインスタンスを指定
#     - タグのキー
#       - env
#     - タグの値
#       - production
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

resource "cloudautomator_job" "example-ec2-start-backup-job" {
  name           = "example-ec2-start-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "ec2_start_backup_job"
  ec2_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    specify_instance            = "tag"
    tag_key                     = "env"
    tag_value                   = "production"
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
