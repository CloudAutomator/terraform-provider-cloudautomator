resource "cloudautomator_job_workflow" "example-job-workflow" {
  name         = "example-job-workflow-basic"
  active       = true
  group_id     = 10
  first_job_id = 1
  follow_job_ids = [
    2,
    3,
  ]
}
