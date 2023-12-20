# ----------------------------------------------------------
# - アクション
#   - ECS: タスクを実行 (Fargate)
# - アクションの設定
#   - リージョン
#     - ap-northeast-1
#   - 対象のECSクラスターの名前
#     - example-cluster
#   - タスクが使用するプラットフォームのバージョン
#     - LATEST
#   - タスク定義のファミリー
#     - example-service
#   - 起動するタスク数
#     - 1
#   - タグをタスク定義からタスクに伝播するかどうか
#     - 伝播させる
#   - タスクにAmazon ECS管理タグを付与するかどうか
#     - 付与する
#   - 使用するVPC
#     - vpc-00000001
#   - 使用するサブネット
#     - subnet-00000001
#     - subnet-00000002
#   - 使用するセキュリティグループ
#     - sg-00000001
#     - sg-00000002
#   - パブリックIP割当を有効にするかどうか
#     - 有効にする
# ----------------------------------------------------------

resource "cloudautomator_job" "example-run-ecs-tasks-fargate-job" {
  name           = "example-run-ecs-tasks-fargate-job"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "run_ecs_tasks_fargate"
  run_ecs_tasks_fargate_action_value {
    region                      = "ap-northeast-1"
    ecs_cluster                 = "example-cluster"
    platform_version            = "LATEST"
    ecs_task_definition_family  = "example-service"
    ecs_task_count              = 1
    propagate_tags              = "TASK_DEFINITION"
    enable_ecs_managed_tags     = true
    ecs_awsvpc_vpc              = "vpc-00000001"
    ecs_awsvpc_subnets          = ["subnet-00000001", "subnet-00000002"]
    ecs_awsvpc_security_groups  = ["sg-00000001", "sg-00000002"]
    ecs_awsvpc_assign_public_ip = "ENABLED"
  }
}
