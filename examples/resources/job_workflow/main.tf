# ----------------------------------------------------------
# - ジョブワークフロー
#  - 名前
#    - example-job-workflow-basic
#  - 状態
#    - 有効
#  - グループID
#    - 10
#  - 先頭ジョブID
#    - 1
#  - 後続ジョブID
#    - 2
#    - 3
# ----------------------------------------------------------

resource "cloudautomator_job_workflow" "job-workflow-basic" {
  name         = "example-job-workflow-basic"
  active       = true
  group_id     = 10
  first_job_id = 1
  follow_job_ids = [
    2,
    3,
  ]
}
