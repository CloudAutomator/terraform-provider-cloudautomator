# ----------------------------------------------------------
# - アクション
#   - EC2: インスタンスでコマンドを実行
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のEC2インスタンスを指定
#     - インスタンス特定に利用するタグのキー
#       - env
#     - インスタンス特定に利用するタグの値
#       - production
#   - 実行するコマンド
#     - whoami
#   - コマンドに設定するコメント
#     - test
#   - コマンドの種類
#     - AWS-RunPowerShellScript
#   - 実行結果を保存するS3のバケット名
#     - example-bucket
#   - 実行結果を保存するS3のプレフィックス
#     - example
#   - 実行コマンドの終了ステータスをジョブ完了の判定にするフラグ
#     - true
#   - インスタンス接続のタイムアウト時間(秒)
#     - 30
#   - コマンド実行のタイムアウト時間(秒)
#     - 30
# ----------------------------------------------------------

resource "cloudautomator_job" "example-send-command-job" {
  name           = "example-send-command-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "send_command"
  send_command_action_value {
    region                    = "ap-northeast-1"
    specify_instance          = "tag"
    tag_key                   = "env"
    tag_value                 = "production"
    command                   = "whoami"
    comment                   = "test"
    document_name             = "AWS-RunShellScript"
    output_s3_bucket_name     = "example-bucket"
    output_s3_key_prefix      = "example"
    trace_status              = "true"
    timeout_seconds           = 30
    execution_timeout_seconds = 30
  }
}
