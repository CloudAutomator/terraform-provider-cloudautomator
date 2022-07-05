# ----------------------------------------------------------
# - アクション
#   - RDS(Aurora): DBクラスタースナップショットからリストア
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - リストア後のDBインスタンスのidentifier
#     - example-db
#   - リストア後のDBクラスターのidentifier
#     - example-cluster
#   - リストアに使用するDBスナップショットID
#     - example-snapshot
#   - リストア後のDBクラスターのDBエンジン
#     - mysql
#   - リストア後のDBクラスターのDBエンジンのバージョン
#     - 1.2.3.4
#   - リストア後のDBクラスターのDBインスタンスクラス
#     - db.t2.micro
#   - リストア後のDBクラスターを配置するDBサブネットグループ名
#     - example-subnet-group
#   - リストア後のDBクラスターをパブリックアクセス可能にするか否か
#     - する
#   - リストア後のDBクラスターを配置するAZ
#     - ap-northeast-1
#   - リストア後のDBクラスターに設定するセキュリティグループIDが含まれる配列
#     - sg-1234567
#   - リストア後のDBクラスターの接続ポート番号
#     - 5432
#   - リストア後のDBクラスターに設定するパラメータグループ名
#     - example-cluster-parameter-group
#   - リストア後のDBインスタンスに設定するパラメータグループ名
#     - example-parameter-group
#   - リストア後のDBクラスターに設定するオプショングループ名
#     - example-option-group
#   - リストア後のDBクラスターで自動マイナーバージョンアップグレードを有効にするかどうか
#     - する
#   - リストアに利用したDBスナップショットを削除するかどうか
#     - する
# ----------------------------------------------------------

resource "cloudautomator_job" "example-restore-rds-cluster-job" {
  name           = "example-restore-rds-cluster-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "restore_rds_cluster"
  restore_rds_cluster_action_value {
    region                          = "ap-northeast-1"
    db_instance_identifier          = "example-db"
    db_cluster_identifier           = "example-cluster"
    snapshot_identifier             = "example-snapshot"
    engine                          = "mysql"
    engine_version                  = "1.2.3.4"
    db_instance_class               = "db.t2.micro"
    db_subnet_group_name            = "example-subnet-group"
    publicly_accessible             = "false"
    availability_zone               = "ap-northeast-1"
    vpc_security_group_ids          = ["sg-1234567"]
    port                            = "5432"
    db_cluster_parameter_group_name = "example-cluster-parameter-group"
    db_parameter_group_name         = "example-parameter-group"
    option_group_name               = "example-option-group"
    auto_minor_version_upgrade      = "true"
    delete_db_cluster_snapshot      = "true"
  }
}
