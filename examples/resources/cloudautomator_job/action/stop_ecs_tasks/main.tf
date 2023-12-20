# ----------------------------------------------------------
# - アクション
#   - ECS: タスクを停止
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - 対象のECSクラスターの名前
#    - example-cluster
#   - ECSタスクの特定方法
#     - タグ
#   - タグのキー
#     - env
#   - タグの値
#     - develop
# ----------------------------------------------------------

resource "cloudautomator_job" "example-stop-ecs-tasks" {
  name           = "example-stop-ecs-tasks"
  group_id       = 10
  aws_account_id = 20

  rule_type = "webhook"

  action_type = "stop_ecs_tasks"
  stop_ecs_tasks_action_value {
    region           = "ap-northeast-1"
    ecs_cluster      = "example-cluster"
    specify_ecs_task = "tag"
    tag_key          = "env"
    tag_value        = "develop"
  }
}
