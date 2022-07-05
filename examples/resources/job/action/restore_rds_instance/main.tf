# ----------------------------------------------------------
# - アクション
#   - RDS: DBスナップショットからリストア
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - リストア後のRDSインスタンス名
#     - example-db
#   - リストアに使用するDBスナップショットID
#     - example-snapshot
#   - リストア後のRDSインスタンスのDBエンジン
#     - mysql
#   - リストア後のRDSインスタンスのライセンスモデル
#     - license-included
#   - リストア後のRDSインスタンスのDBインスタンスクラス
#     - db.t2.micro
#   - リストア後のRDSインスタンスをMulti-AZ構成にするか否か
#     - する
#   - リストア後のRDSインスタンスのストレージタイプ
#     - standard
#   - リストア後のRDSインスタンスのIOPS値
#     - 1000
#   - リストア後のRDSインスタンスを配置するVPCのID
#     - vpc-00000001”
#   - リストア後のRDSインスタンスを配置するDBサブネットグループ名
#     - example-subnet-group
#   - リストア後のRDSインスタンスをパブリックアクセス可能にするか否か
#     - する
#   - リストア後のRDSインスタンスを配置するAZ
#     - ap-northeast-1
#   - リストア後のRDSインスタンスに設定するセキュリティグループIDが含まれる配列
#     - sg-1234567
#   - リストア後のRDSインスタンスのデータベース名
#     - example-db
#   - リストア後のRDSインスタンスの接続ポート番号
#     - 5432
#   - リストア後のRDSインスタンスに設定するパラメータグループ名
#     - example-parameter-group
#   - リストア後のRDSインスタンスに設定するオプショングループ名
#     - example-option-group
#   - リストア後のRDSインスタンスで自動マイナーバージョンアップグレードを有効にするかどうか
#     - する
#   - リストアに利用したDBスナップショットを削除するかどうか
#     - する
#   - リストア後のRDSインスタンスに割り当てるタグのキー
#     - env
#   - リストア後のRDSインスタンスの割り当てるタグの値
#     - production
#   - RDSインスタンスの作成完了をジョブ完了の判定にするフラグ
#     - する
# ----------------------------------------------------------

resource "cloudautomator_job" "example-restore-rds-instance-job" {
  name           = "example-restore-rds-instance-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "restore_rds_instance"
  restore_rds_instance_action_value {
    region                     = "ap-northeast-1"
    rds_instance_id            = "example-db"
    rds_snapshot_id            = "example-snapshot"
    db_engine                  = "mysql"
    license_model              = "license-included"
    db_instance_class          = "db.t2.micro"
    multi_az                   = "true"
    storage_type               = "standard"
    iops                       = 1000
    vpc                        = "vpc-00000001"
    subnet_group               = "example-subnet-group"
    publicly_accessible        = "false"
    availability_zone          = "ap-northeast-1"
    vpc_security_group_ids     = ["sg-1234567"]
    db_name                    = "example-db"
    port                       = 5432
    parameter_group            = "example-parameter-group"
    option_group               = "example-option-group"
    auto_minor_version_upgrade = "true"
    delete_rds_snapshot        = "true"
    additional_tag_key         = "env"
    additional_tag_value       = "production"
    trace_status               = "true"
  }
}
