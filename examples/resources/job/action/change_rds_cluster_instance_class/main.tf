# ----------------------------------------------------------
# - アクション
#   - RDS(Aurora): DBインスタンスクラスを変更
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のDBインスタンス(Auroraエンジン)を指定
#     - タグのキー
#     - env
#   - タグの値
#     - production
#   - DBインスタンスクラス
#     - db.t3.micro
# ----------------------------------------------------------
resource "cloudautomator_job" "example-change-rds-cluster-instance-class-job" {
  name = "example-change-rds-cluster-instance-class-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "change_rds_cluster_instance_class"
  change_rds_cluster_instance_class_action_value {
    region = "ap-northeast-1"
    specify_rds_instance = "tag"
    tag_key = "env"
    tag_value = "develop"
    db_instance_class = "db.t3.micro"
  }
}
