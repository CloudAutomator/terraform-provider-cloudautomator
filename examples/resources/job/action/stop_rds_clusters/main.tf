# ----------------------------------------------------------
# - アクション
#   - RDS(Aurora): DBクラスターを停止
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のDBクラスターを指定
#     - DBクラスター特定に利用するタグのキー
#       - env
#     - DBクラスター特定に利用するタグの値
#       - production
#   - DBクラスターの停止完了をジョブ完了の判定にするフラグ
#     - true
# ----------------------------------------------------------

resource "cloudautomator_job" "example-stop-rds-clusters-job" {
  name           = "example-stop-rds-clusters-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "stop_rds_clusters"
  stop_rds_clusters_action_value {
    region              = "ap-northeast-1"
    specify_rds_cluster = "tag"
    tag_key             = "env"
    tag_value           = "production"
    trace_status        = "true"
  }
}
