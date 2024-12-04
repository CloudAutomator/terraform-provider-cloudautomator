# ----------------------------------------------------------
# - アクション
#   - EC2: セキュリティグループにインバウンドルールを追加
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のセキュリティグループを指定
#     - タグのキー
#       - env
#     - タグの値
#       - production
#   - 通信プロトコル
#     - TCP
#   - ポート番号
#     - 80
#   - 送信元IPのCIDRアドレス
#     - 172.31.0.0/16
# ----------------------------------------------------------

resource "cloudautomator_job" "example-authorize-security-group-ingress-job" {
  name           = "example-authorize-security-group-ingress-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "authorize_security_group_ingress"
  authorize_security_group_ingress_action_value {
    region                 = "ap-northeast-1"
    specify_security_group = "tag"
    tag_key                = "env"
    tag_value              = "develop"
    ip_protocol            = "tcp"
    to_port                = "80"
    cidr_ip                = "172.31.0.0/16"
  }
}
