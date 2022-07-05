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
#   - Windows Updateの適用で発生する再起動を許容するか
#     - しない
#   - 適用するWindows Updateの重要度
#     - Critical, Important
#   - 実行結果を保存するS3のバケット名
#     - example-bucket
#   - 実行結果を保存するS3のプレフィックス
#     - example-prefix
#   - Windows Update完了をジョブ完了の判定にするかどうか
#     - する
# ----------------------------------------------------------

resource "cloudautomator_job" "example-windows-update-v2-job" {
  name           = "example-windows-update-v2-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "windows_update_v2"
  windows_update_v2_action_value {
    region                = "ap-northeast-1"
    specify_instance      = "tag"
    tag_key               = "env"
    tag_value             = "production"
    allow_reboot          = "false"
    severity_levels       = ["Critical", "Important"]
    output_s3_bucket_name = "example-bucket"
    output_s3_key_prefix  = "example-prefix"
    trace_status          = "true"
  }
}
