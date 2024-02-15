# ----------------------------------------------------------
# - アクション
#   - DynamoDB: テーブルをバックアップ
# - アクションの設定
#   - 対象のテーブルとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のDynamoDBテーブルの名前
#     - example-table
#   - 対象のバックアップボールトの名前
#     - example-backup
#   - バックアップの保持期間(日数)
#     - 無期限
#   - バックアップ取得時に使うIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForBackup
# ----------------------------------------------------------

resource "cloudautomator_job" "example-dynamodb-start-backup-job" {
  name           = "example-dynamodb-start-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "dynamodb_start_backup_job"
  dynamodb_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    dynamodb_table_name         = "example-table"
    backup_vault_name           = "example-backup"
    lifecycle_delete_after_days = null
    iam_role_arn                = "arn:aws:iam::123456789012:role/RoleForBackup"
  }
}

# ----------------------------------------------------------
# - アクション
#   - DynamoDB: テーブルをバックアップ
# - アクションの設定
#   - 対象のテーブルとバックアップボールトが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のDynamoDBテーブルの名前
#     - example-table
#   - バックアップの保持期間(日数)
#     - 7
#   - 対象のバックアップボールトの名前
#     - example-backup
#   - バックアップ取得時に使うIAMロールのARN
#     - arn:aws:iam::123456789012:role/RoleForBackup
#   - 作成した復旧ポイントに割り当てるタグの配列
#     - key1: value1
#     - key2: value2
#     - key3: value3
# ----------------------------------------------------------

resource "cloudautomator_job" "example-dynamodb-start-backup-job" {
  name           = "example-dynamodb-start-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "dynamodb_start_backup_job"
  dynamodb_start_backup_job_action_value {
    region                      = "ap-northeast-1"
    dynamodb_table_name         = "example-table"
    lifecycle_delete_after_days = 7
    backup_vault_name           = "example-backup"
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
