# ----------------------------------------------------------
# - アクション: ElastiCache: ノードタイプを変更
# - アクションの設定:
#   - 変更対象のElastiCacheクラスターが存在するAWSリージョン
#     - ap-northeast-1
#   - ElastiCacheクラスター特定に利用するタグ
#     - キー: env
#     - 値: production
#   - 変更後のノードタイプ
#     - cache.t4g.small
# ----------------------------------------------------------

resource "cloudautomator_job" "example-change-elasticache-node-type" {
  name           = "example-change-elasticache-node-type"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "change_elasticache_node_type"
  change_elasticache_node_type_action_value {
    region    = "ap-northeast-1"
    tag_key   = "env"
    tag_value = "production"
    node_type = "cache.t4g.small"
  }
}
