# ----------------------------------------------------------
# - アクション
#   - EC2: セキュリティグループからインバウンドルールを削除
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - タグで対象のセキュリティグループを指定
#     - セキュリティグループ特定に利用するタグのキー
#       - env
#     - セキュリティグループ特定に利用するタグの値
#       - production
#   - 通信プロトコル
#     - tcp
#   - ポート番号
#     - 80
#   - 送信元IPのCIDRアドレス
#     - 172.31.0.0/16
# ----------------------------------------------------------

resource "cloudautomator_job" "example-revoke-security-group-ingress-job" {
  name           = "example-revoke-security-group-ingress-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "revoke_security_group_ingress"
  revoke_security_group_ingress_action_value {
    region                 = "ap-northeast-1"
    specify_security_group = "tag"
    tag_key                = "env"
    tag_value              = "production"
    ip_protocol            = "tcp"
    to_port                = "80"
    cidr_ip                = "172.31.0.0/16"
  }
}
