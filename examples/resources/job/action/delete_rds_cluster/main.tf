# ----------------------------------------------------------
# - アクション
#   - RDS(Aurora): DBクラスターを削除
# - アクションの設定
#   - DBクラスターを削除するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のDBクラスターを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - DBクラスター削除時に取得するDBクラスタースナップショット名
#     - test-snapshot
#   - DBクラスター削除時のDBクラスタースナップショット取得をスキップする
#     - false
#   - DBクラスターの削除完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------
resource "cloudautomator_job" "example-delete-rds-cluster-job" {
  name = "example-delete-rds-cluster-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "delete_rds_cluster"
  delete_rds_cluster_action_value {
    region = "ap-northeast-1"
    specify_rds_cluster = "tag"
    tag_key = "env"
    tag_value = "develop"
    final_db_snapshot_identifier = "test-snapshot"
    skip_final_snapshot = "false"
    trace_status = "true"
  }
}
