# ----------------------------------------------------------
# - Webhook後処理
#   - グループID
#     - 123
#   - Authorizationヘッダの値
#     - test-authorization-header
#   - Webhook送信先となるURL
#     - https://example.com/webhook
# ----------------------------------------------------------
resource "cloudautomator_post_process" "webhook" {
  name = "webhook"
  group_id = 10
  service = "webhook"

  webhook_parameters {
    webhook_authorization_header = "test-authorization-header"
    webhook_url = "http://example.com"
  }
}
