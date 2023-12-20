# ----------------------------------------------------------
# - アクション
#   - Compute Engine: マシンイメージを作成
# - アクションの設定
#   - リージョン
#     - asia-northeast1
#   - プロジェクトID
#     - example-project
#   - VMインスタンスをラベルで指定
#     - ラベルのキー
#       - env
#     - ラベルの値
#       - develop
#   - マシンイメージの保存場所
#     - asia-northeast1
#   - 作成するマシンイメージの名前
#     - example-daily
#   - マシンイメージの世代管理を行う数
#     - 10
# ----------------------------------------------------------

resource "cloudautomator_job" "example-google-compute-insert-machine-image" {
  name                    = "example-google-compute-insert-machine-image"
  group_id                = 10
  google_cloud_account_id = 20

  rule_type = "webhook"

  action_type = "google_compute_insert_machine_image"
  google_compute_insert_machine_image_action_value {
    region                         = "asia-northeast1"
    project_id                     = "example-project"
    specify_vm_instance            = "label"
    vm_instance_label_key          = "env"
    vm_instance_label_value        = "develop"
    machine_image_storage_location = "asia-northeast1"
    machine_image_basename         = "example-daily"
    generation                     = 10
  }
}
