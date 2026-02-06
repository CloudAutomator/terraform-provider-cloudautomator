# ----------------------------------------------------------
# - アクション: RDS(Aurora): DBクラスターのACUを変更
# - アクションの設定:
#   - 対象リージョン: ap-northeast-1
#   - DBクラスター特定方法: タグ
#   - タグキー: env
#   - タグ値: production
#   - 最小容量: 0.5 ACU
#   - 最大容量: 1 ACU
#   - 自動停止までの待機時間: 300秒
# ----------------------------------------------------------

resource "cloudautomator_job" "example-rds_cluster_change_capacity" {
  name           = "example-rds_cluster_change_capacity"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "rds_cluster_change_capacity"
  rds_cluster_change_capacity_action_value {
    region                   = "ap-northeast-1"
    specify_rds_cluster      = "tag"
    tag_key                  = "env"
    tag_value                = "production"
    acu_min_capacity         = 0.5
    acu_max_capacity         = 1
    seconds_until_auto_pause = 300
  }

  completed_post_process_id = [1, 2]
  failed_post_process_id    = [1, 2]
}
