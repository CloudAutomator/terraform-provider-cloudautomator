# ----------------------------------------------------------
# - アクション
#   - ECS: サービスのタスク数を変更
# - アクションの設定
#   - 対象のECSクラスターおよびECSサービスが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のECSクラスター
#     - test-cluster
#   - 対象のECSサービスの特定方法
#     - サービス名で特定
#   - 対象のECSサービス
#     - test-service
#   - タスク数変更方法
#     - DesiredCount の直接変更
#   - ECSサービスのタスク数（DesiredCount）
#     - 3
# ----------------------------------------------------------

resource "cloudautomator_job" "example-ecs-change-service-task-count" {
  name           = "example-ecs-change-service-task-count"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "ecs_change_service_task_count"
  ecs_change_service_task_count_action_value {
    region              = "ap-northeast-1"
    ecs_cluster         = "test-cluster"
    specify_ecs_service = "identifier"
    ecs_service         = "test-service"
    specify_task_change = "task"
    desired_count       = 3
  }
}

# ----------------------------------------------------------
# - アクション
#   - ECS: サービスのタスク数を変更
# - アクションの設定
#   - 対象のECSクラスターおよびECSサービスが存在するAWSリージョン
#     - ap-northeast-1
#   - 対象のECSクラスター
#     - test-cluster
#   - 対象のECSサービスの特定方法
#     - タグで特定
#   - サービス特定に利用するタグ - Key
#     - env
#   - サービス特定に利用するタグ - Value
#     - production
#   - タスク数変更方法
#     - AutoScaling の MinCapacity/MaxCapacity 変更
#   - AutoScaling 設定のタスクの最小数（MinCapacity）
#     - 2
#   - AutoScaling 設定のタスクの最大数（MaxCapacity）
#     - 10
# ----------------------------------------------------------

resource "cloudautomator_job" "example-ecs-change-service-task-count-autoscaling" {
  name           = "example-ecs-change-service-task-count-autoscaling"
  group_id       = 10
  aws_account_id = 20

  rule_type = "immediate_execution"

  action_type = "ecs_change_service_task_count"
  ecs_change_service_task_count_action_value {
    region              = "ap-northeast-1"
    ecs_cluster         = "test-cluster"
    specify_ecs_service = "tag"
    tag_key             = "env"
    tag_value           = "production"
    specify_task_change = "autoscaling"
    min_capacity        = 2
    max_capacity        = 10
  }
}
