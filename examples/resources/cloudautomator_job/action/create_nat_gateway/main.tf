# ----------------------------------------------------------
# - アクション: VPC: NAT Gatewayを作成
# - アクションの設定:
#   - NAT Gatewayを作成するAWSリージョン
#     - ap-northeast-1
#   - NAT Gatewayに割り当てるElasticIPの割当ID
#     - eipalloc-0123456789abcdef0
#   - NAT GatewayのNameタグに設定する値
#     - Dev-NATGateway
#   - NAT Gatewayを作成するサブネットのID
#     - subnet-0123456789abcdef0
#   - NAT Gatewayを宛先とするルートを追加するルートテーブルのID
#     - rtb-0123456789abcdef0
#   - NAT Gatewayに割り当てるタグの配列
#     - Environment: Development
#     - Project: WebApp
#     - Owner: DevOps
# ----------------------------------------------------------

resource "cloudautomator_job" "example-create-nat-gateway" {
  name           = "example-create-nat-gateway"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "create_nat_gateway"
  create_nat_gateway_action_value {
    region           = "ap-northeast-1"
    allocation_id    = "eipalloc-0123456789abcdef0"
    nat_gateway_name = "Dev-NATGateway"
    subnet_id        = "subnet-0123456789abcdef0"
    route_table_id   = "rtb-0123456789abcdef0"

    additional_tags {
      key   = "Environment"
      value = "Development"
    }

    additional_tags {
      key   = "Project"
      value = "WebApp"
    }

    additional_tags {
      key   = "Owner"
      value = "DevOps"
    }
  }
}