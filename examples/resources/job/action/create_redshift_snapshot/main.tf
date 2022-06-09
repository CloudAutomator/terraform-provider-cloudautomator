# ----------------------------------------------------------
# - アクション
#   - Redshift: クラスタースナップショットを作成
# - アクションの設定
#   - スナップショットを作成するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のRedshiftクラスターを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - スナップショットに設定する名前
#     - test
#   - スナップショットの世代管理を行う数
#     - 10
#   - スナップショットの作成完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------
resource "cloudautomator_job" "example-create-redshift-snapshot-job" {
  name = "example-create-redshift-snapshot-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "create_redshift_snapshot"
  create_redshift_snapshot_action_value {
    region = "ap-northeast-1"
    specify_cluster = "tag"
    tag_key = "env"
    tag_value = "develop"
    generation = 10
    cluster_snapshot_identifier = "test"
    trace_status = "true"
  }
}
