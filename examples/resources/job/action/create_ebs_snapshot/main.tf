# ----------------------------------------------------------
# - アクション
#   - EC2: EBSスナップショットを作成
# - アクションの設定
#   - EBSスナップショットを作成するAWSリージョン
#     - ap-northeast-1
#   - タグで対象のEBSボリュームを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - EBSボリュームの世代管理を行う数
#     - 10
#   - EBSボリュームに設定する説明
#     - test db
#   - 作成したEBSボリュームに割り当てるタグのキー
#     - example-key
#   - 作成したEBSボリュームに割り当てるタグの値
#     - example-value
#   - EBSボリュームの作成完了をジョブ完了の判定にする
#     - true
# ----------------------------------------------------------
resource "cloudautomator_job" "example-create-ebs-snapshot-job" {
  name = "example-create-ebs-snapshot-job"
  group_id = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "create_ebs_snapshot"
  create_ebs_snapshot_action_value {
    region = "ap-northeast-1"
    specify_volume = "tag"
    tag_key = "env"
    tag_value = "develop"
    generation = 10
    description = "test db"
    additional_tag_key = "example-key"
    additional_tag_value = "example-value"
    trace_status = "true"
  }
}
