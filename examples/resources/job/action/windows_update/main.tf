# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスをWindows Update
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のEC2インスタンスを指定
#     - EC2インスタンス特定に利用するタグのキー
#       - env
#     - EC2インスタンス特定に利用するタグの値
#       - production
#   - コマンドに設定するコメント
#     - example-comment
#   - ドキュメント名
#     - AWS-InstallMissingWindowsUpdates
#   - 除外するKBが含まれた配列
#     - KB1111111
#   - 実行結果を保存するS3のバケット名
#     - example-bucket
#   - 実行結果を保存するS3のプレフィックス
#     - example-prefix
#   - アップデートレベル
#     - All
#   - タイムアウト時間(秒)
#     - 3600
# ----------------------------------------------------------

resource "cloudautomator_job" "example-windows-update-job" {
  name           = "example-windows-update-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "windows_update"
  windows_update_action_value {
    region                = "ap-northeast-1"
    specify_instance      = "tag"
    tag_key               = "env"
    tag_value             = "production"
    comment               = "example-comment"
    document_name         = "AWS-InstallMissingWindowsUpdates"
    kb_article_ids        = ["KB1111111"]
    output_s3_bucket_name = "example-bucket"
    output_s3_key_prefix  = "example-prefix"
    update_level          = "All"
    timeout_seconds       = "3600"
  }
}
