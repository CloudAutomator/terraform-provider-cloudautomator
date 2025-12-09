package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Ecs(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "RunEcsTasksFargateAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "run_ecs_tasks_fargate"
				run_ecs_tasks_fargate_action_value {
					region = "ap-northeast-1"
					ecs_cluster = "example-cluster"
					platform_version = "LATEST"
					ecs_task_definition_family = "example-service"
					ecs_task_count = 1
					propagate_tags = "TASK_DEFINITION"
					enable_ecs_managed_tags = true
					ecs_awsvpc_vpc = "vpc-00000001"
					ecs_awsvpc_subnets = ["subnet-00000001", "subnet-00000002"]
					ecs_awsvpc_security_groups = ["sg-00000001", "sg-00000002"]
					ecs_awsvpc_assign_public_ip = "ENABLED"
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "run_ecs_tasks_fargate"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_cluster", "example-cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.platform_version", "LATEST"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_task_definition_family", "example-service"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_task_count", "1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.propagate_tags", "TASK_DEFINITION"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.enable_ecs_managed_tags", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_vpc", "vpc-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_subnets.#", "2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_subnets.0", "subnet-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_subnets.1", "subnet-00000002"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_security_groups.#", "2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_security_groups.0", "sg-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_security_groups.1", "sg-00000002"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "run_ecs_tasks_fargate_action_value.0.ecs_awsvpc_assign_public_ip", "ENABLED"),
			},
		},
		{
			name:    "StopEcsTasksAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "stop_ecs_tasks"
				stop_ecs_tasks_action_value {
					region = "ap-northeast-1"
					ecs_cluster = "example-cluster"
					specify_ecs_task = "tag"
					tag_key = "env"
					tag_value = "develop"
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "stop_ecs_tasks"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_ecs_tasks_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_ecs_tasks_action_value.0.ecs_cluster", "example-cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_ecs_tasks_action_value.0.specify_ecs_task", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_ecs_tasks_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_ecs_tasks_action_value.0.tag_value", "develop"),
			},
		},
		{
			name:    "EcsChangeServiceTaskCountAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "ecs_change_service_task_count"
				ecs_change_service_task_count_action_value {
					region = "ap-northeast-1"
					ecs_cluster = "example-cluster"
					specify_ecs_service = "tag"
					tag_key = "env"
					tag_value = "production"
					specify_task_change = "task"
					desired_count = 3
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "ecs_change_service_task_count"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.ecs_cluster", "example-cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.specify_ecs_service", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.tag_value", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.specify_task_change", "task"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ecs_change_service_task_count_action_value.0.desired_count", "3"),
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			baseChecks := []resource.TestCheckFunc{
				testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], "cloudautomator_job.test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "name", tc.jobName),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "completed_post_process_id.0", acctest.TestPostProcessId()),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "failed_post_process_id.0", acctest.TestPostProcessId()),
			}

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(tc.jobName),
						Check: resource.ComposeAggregateTestCheckFunc(
							append(baseChecks, tc.checks...)...,
						),
					},
				},
			})
		})
	}
}
