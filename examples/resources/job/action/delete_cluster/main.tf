# ----------------------------------------------------------
# - アクション
#   - Redshift: クラスターを削除
# - アクションの設定
#   - Redshiftクラスターを削除するAWSリージョン
#     - ap-northeast-1
#   - 対象のRedshiftクラスターのidentifier
#     - test-cluster
#   - Redshiftクラスター削除時に取得するRedshiftクラスタースナップショット名
#     - test-snapshot
#   - Redshiftクラスター削除時のRedshiftクラスタースナップショット取得をスキップする
#     - false
#   - Redshiftクラスターの削除完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------

resource "cloudautomator_job" "example-delete-cluster-job" {
  name           = "example-delete-cluster-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "delete_cluster"
  delete_cluster_action_value {
    region                            = "ap-northeast-1"
    cluster_identifier                = "test-cluster"
    final_cluster_snapshot_identifier = "test-snapshot"
    skip_final_cluster_snapshot       = "false"
    trace_status                      = "true"
  }
}
