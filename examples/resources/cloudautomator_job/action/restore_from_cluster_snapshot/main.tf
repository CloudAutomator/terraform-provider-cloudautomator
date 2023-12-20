# ----------------------------------------------------------
# - アクション
#   - Redshift: スナップショットからリストア
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - リストア後のRedshiftクラスターのidentifier
#     - example-cluster
#   - リストアに使用するRedshiftスナップショットID
#     - example-snapshot
#   - リストア後のRedshiftクラスターに設定するパラメータグループ名
#     - example-parameter-group
#   - リストア後のRedshiftクラスターを配置するサブネットグループ名
#     - exmaple-subnet-group
#   - リストア後のDBクラスターの接続ポート番号
#     - 5432
#   - リストア後のRedshiftクラスターをパブリックアクセス可能にするか否か
#     - 不可
#   - リストア後のRedshiftクラスターに設定するセキュリティグループIDが含まれる配列
#     - sg-1234567
#   - リストア後のRedshiftクラスターで自動マイナーバージョンアップグレードを有効にするかどうか
#     - する
#   - リストアに利用したRedshiftスナップショットを削除するかどうか
#     - する
# ----------------------------------------------------------

resource "cloudautomator_job" "example-restore-from-cluster-snapshot-job" {
  name           = "example-restore-from-cluster-snapshot-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "restore_from_cluster_snapshot"
  restore_from_cluster_snapshot_action_value {
    region                       = "ap-northeast-1"
    cluster_identifier           = "example-cluster"
    snapshot_identifier          = "example-snapshot"
    cluster_parameter_group_name = "example-parameter-group"
    cluster_subnet_group_name    = "example-subnet-group"
    port                         = "5432"
    publicly_accessible          = "false"
    vpc_security_group_ids       = ["sg-1234567"]
    allow_version_upgrade        = "true"
    delete_cluster_snapshot      = "true"
  }
}
