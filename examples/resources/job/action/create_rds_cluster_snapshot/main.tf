# ----------------------------------------------------------
# - アクション
#   - RDS: DBスナップショットを作成
# - アクションの設定
#   - DBスナップショットを作成するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のDBインスタンスを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - DBスナップショットに設定する名前
#     - test
#   - DBスナップショットの世代管理を行う数
#     - 10
#   - DBスナップショットの作成完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------

resource "cloudautomator_job" "example-create-rds-snapshot-job" {
  name           = "example-create-rds-snapshot-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "create_rds_cluster_snapshot"
  create_rds_cluster_snapshot_action_value {
    region                         = "ap-northeast-1"
    specify_rds_cluster            = "tag"
    tag_key                        = "env"
    tag_value                      = "develop"
    generation                     = 10
    db_cluster_snapshot_identifier = "test"
    trace_status                   = "true"
  }
}
