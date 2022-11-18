# ----------------------------------------------------------
# - アクション
#   - FSx: バックアップを作成
# - アクションの設定
#   - バックアップを作成するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のファイルシステムを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - バックアップの世代管理を行う数
#     - 10
#   - バックアップ名
#     - example-backup
# ----------------------------------------------------------

resource "cloudautomator_job" "example-create-fsx-backup-job" {
  name           = "example-create-fsx-backup-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "create_fsx_backup"
  create_fsx_backup_action_value {
    region              = "ap-northeast-1"
    specify_file_system = "tag"
    tag_key             = "env"
    tag_value           = "develop"
    generation          = 10
    backup_name         = "example-backup"
  }
}
