# ----------------------------------------------------------
# - アクション
#   - Compute Engine: VMインスタンスを停止
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
# ----------------------------------------------------------

resource "cloudautomator_job" "example-google-compute-stop-vm-instances-image" {
  name                    = "example-google-compute-stop-vm-instances-image"
  group_id                = 10
  google_cloud_account_id = 20

  rule_type = "webhook"

  action_type = "google_compute_stop_vm_instances"
  google_compute_stop_vm_instances_action_value {
    region                  = "asia-northeast1"
    project_id              = "example-project"
    specify_vm_instance     = "label"
    vm_instance_label_key   = "env"
    vm_instance_label_value = "develop"
  }
}
