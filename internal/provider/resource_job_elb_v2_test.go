package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_ElbV2(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "RegisterTargetInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "register_target_instances"
				register_target_instances_action_value {
					region = "ap-northeast-1"
					target_group_arn = "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/t1/c8a1987f0402f55a"
					tag_key = "env"
					tag_value = "develop"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "register_target_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_target_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_target_instances_action_value.0.target_group_arn", "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/t1/c8a1987f0402f55a"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_target_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_target_instances_action_value.0.tag_value", "develop"),
			},
		},
		{
			name:    "DeregisterTargetInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "deregister_target_instances"
				deregister_target_instances_action_value {
					region = "ap-northeast-1"
					target_group_arn = "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/t1/c8a1987f0402f55a"
					tag_key = "env"
					tag_value = "develop"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "deregister_target_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_target_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_target_instances_action_value.0.target_group_arn", "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/t1/c8a1987f0402f55a"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_target_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_target_instances_action_value.0.tag_value", "develop"),
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
