# ----------------------------------------------------------
# - アクション
#   - RDS(Aurora): DBクラスタースナップショットをリージョン間でコピー
# - アクションの設定
#   - DBクラスタースナップショットのコピー元のAWSリージョン
#     - ap-northeast-1
#   - DBクラスタースナップショットのコピー先のAWSリージョン
#     - ap-southeast-1
#   - IDで対象のDBクラスタースナップショットを指定
#     - 対象のDBスナップショットID
#       - rds_cluster_snapshot_id
#   - KMSキーID
#     - 1234abcd-12ab-34cd-56ef-1234567890ab
# ----------------------------------------------------------


resource "cloudautomator_job" "example-copy-rds-cluster-snapshot-job" {
  name = "example-copy-rds-cluster-snapshot-job"

  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "copy_rds_cluster_snapshot"
  copy_rds_cluster_snapshot_action_value {
    source_region                = "ap-northeast-1"
    destination_region           = "ap-southeast-1"
    specify_rds_cluster_snapshot = "rds_cluster_snapshot_id"
    rds_cluster_snapshot_id      = "test-snapshot"
    kms_key_id                   = "1234abcd-12ab-34cd-56ef-1234567890ab"
  }
}
