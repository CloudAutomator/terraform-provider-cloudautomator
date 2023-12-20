# ----------------------------------------------------------
# - アクション
#   - RDS(Aurora): DBクラスターを起動
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のDBクラスターを指定
#     - DBクラスター特定に利用するタグのキー
#       - env
#     - DBクラスター特定に利用するタグの値
#       - production
#   - DBクラスターの起動完了をジョブ完了の判定にするフラグ
#     - true
# ----------------------------------------------------------

resource "cloudautomator_job" "example-start-rds-clusters-job" {
  name           = "example-start-rds-clusters-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "start_rds_clusters"
  start_rds_clusters_action_value {
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
    timeout_seconds           = "30"
    execution_timeout_seconds = "30"
  }
}
