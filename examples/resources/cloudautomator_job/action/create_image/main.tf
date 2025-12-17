# ----------------------------------------------------------
# - アクション
#   - EC2: AMIを作成
# - アクションの設定
#   - AMIを作成するAWSリージョン
#     - ap-northeast-1
#   - IDで対象のEC2インスタンスを指定
#     - タグのキー
#       - env
#     - タグの値
#       - develop
#   - AMIの世代管理を行う数
#     - 10
#   - AMIに設定するイメージ名
#     - test-image
#   - AMIに設定する説明
#     - test image
#   - AMI 作成時にインスタンスを再起動するか否か
#     - true
#   - 作成した AMI に割り当てるタグの配列
#     - 1つ目のタグ
#       - タグのキー
#         - key-1
#       - タグの値
#         - value-1
#     - 2つ目のタグ
#       - タグのキー
#         - key-2
#       - タグの値
#         - value-2
#   - AMIに割り当てたタグをEBSスナップショットにも追加する
#     - true
#   - ジョブ失敗時にリトライを行うか
#     - true
# ----------------------------------------------------------

resource "cloudautomator_job" "example-create-image-job" {
  name           = "example-create-image-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "create_image"
  create_image_action_value {
    region                 = "ap-northeast-1"
    specify_image_instance = "tag"
    tag_key                = "env"
    tag_value              = "develop"
    generation             = 10
    image_name             = "test-image"
    description            = "test image"
    reboot_instance        = "true"
    additional_tags {
      key   = "key-1"
      value = "value-1"
    }
    additional_tags {
      key   = "key-2"
      value = "value-2"
    }
    add_same_tag_to_snapshot            = "true"
    recreate_image_if_ami_status_failed = "true"
  }
}
