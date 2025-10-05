# ----------------------------------------------------------
# - アクション: VPC: NAT Gatewayを削除
# - アクションの設定:
#   - 削除するNAT GatewayのあるAWSリージョン
#     - ap-northeast-1
#   - 削除するNAT Gatewayのタグキー
#     - Name
#   - 削除するNAT Gatewayのタグの値
#     - test-nat-gateway
# ----------------------------------------------------------

resource "cloudautomator_job" "example-delete-nat-gateway" {
  name           = "example-delete-nat-gateway"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "delete_nat_gateway"
  delete_nat_gateway_action_value {
    region    = "ap-northeast-1"
    tag_key   = "Name"
    tag_value = "test-nat-gateway"
  }
}
