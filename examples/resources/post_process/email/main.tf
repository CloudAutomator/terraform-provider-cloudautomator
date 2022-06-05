variable "group_id" {
  default = 123
}

# ----------------------------------------------------------
# - メール後処理
#   - グループID
#     - 123
#   - メールの送信先
#     - test@example.com
# ----------------------------------------------------------
resource "cloudautomator_post_process" "email" {
  name = "email"
  group_id = var.group_id
  service = "email"

  email_parameters {
    email_recipient = "test@example.com"
  }
}
