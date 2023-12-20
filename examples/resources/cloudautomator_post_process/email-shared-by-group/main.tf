# ----------------------------------------------------------
# - メール後処理
#   - グループ共通 (全てのグループで利用できる後処理)
#   - メールの送信先
#     - test@example.com
# ----------------------------------------------------------

resource "cloudautomator_post_process" "email-shared-by-group" {
  name            = "email-shared-by-group"
  shared_by_group = true
  service         = "email"

  email_parameters {
    email_recipient = "test@example.com"
  }
}
