package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Ec2(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "StartInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "start_instances"
				start_instances_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
					status_checks_enable = "true"
				}

				allow_runtime_action_values = false

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "start_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_instances_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_instances_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_instances_action_value.0.trace_status", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_instances_action_value.0.status_checks_enable", "true"),
			},
		},
		{
			name:    "StopInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "stop_instances"
				stop_instances_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}

				allow_runtime_action_values = false

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "stop_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_instances_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_instances_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_instances_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "BulkStopInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_ids = [%s]

				rule_type = "webhook"

				action_type = "bulk_stop_instances"
				bulk_stop_instances_action_value {
					exclude_by_tag = true
					exclude_by_tag_key = "env"
					exclude_by_tag_value = "production"
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "bulk_stop_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_stop_instances_action_value.0.exclude_by_tag", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_stop_instances_action_value.0.exclude_by_tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_stop_instances_action_value.0.exclude_by_tag_value", "production"),
			},
		},
		{
			name:    "ChangeInstanceTypeAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "change_instance_type"
				change_instance_type_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					instance_type = "t2.medium"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "change_instance_type"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_instance_type_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_instance_type_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_instance_type_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_instance_type_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_instance_type_action_value.0.instance_type", "t2.medium"),
			},
		},
		{
			name:    "SendCommandAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "send_command"
				send_command_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					command = "whoami"
					comment = "test"
					document_name = "AWS-RunShellScript"
					output_s3_bucket_name = "test-s3-bucket"
					output_s3_key_prefix = "test-key"
					trace_status = "true"
					timeout_seconds = "60"
					execution_timeout_seconds = "60"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "send_command"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.command", "whoami"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.comment", "test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.document_name", "AWS-RunShellScript"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.output_s3_bucket_name", "test-s3-bucket"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.output_s3_key_prefix", "test-key"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.trace_status", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.timeout_seconds", "60"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "send_command_action_value.0.execution_timeout_seconds", "60"),
			},
		},
		{
			name:    "WindowsUpdateAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "windows_update"
				windows_update_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					comment = "test"
					document_name = "AWS-InstallMissingWindowsUpdates"
					kb_article_ids = "KB1111111,KB2222222"
					output_s3_bucket_name = "test-s3-bucket"
					output_s3_key_prefix = "test-key"
					update_level = "All"
					timeout_seconds = "60"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "windows_update"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.comment", "test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.document_name", "AWS-InstallMissingWindowsUpdates"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.kb_article_ids", "KB1111111,KB2222222"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.output_s3_bucket_name", "test-s3-bucket"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.output_s3_key_prefix", "test-key"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.update_level", "All"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_action_value.0.timeout_seconds", "60"),
			},
		},
		{
			name:    "WindowsUpdateV2Action",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "windows_update_v2"
				windows_update_v2_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					allow_reboot = "true"
					specify_severity = "select"
					severity_levels = [
						"Critical",
						"Low"
					]
					output_s3_bucket_name = "test-s3-bucket"
					output_s3_key_prefix = "test-key"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "windows_update_v2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.allow_reboot", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.severity_levels.0", "Critical"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.severity_levels.1", "Low"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.output_s3_bucket_name", "test-s3-bucket"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.output_s3_key_prefix", "test-key"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "windows_update_v2_action_value.0.trace_status", "true"),
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
