resource "cloudautomator_post_process" "example-email-post-process" {
  name     = "example-email-post-process"
  group_id = 10
  service  = "email"

  email_parameters {
    email_recipient = "test@example.com"
  }
}
