package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Dynamodb(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "DynamodbStartBackupJobAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "dynamodb_start_backup_job"
				dynamodb_start_backup_job_action_value {
				  region                      = "ap-northeast-1"
				  dynamodb_table_name         = "example-table"
				  lifecycle_delete_after_days = 7
				  backup_vault_name           = "example-vault"
				  iam_role_arn                = "arn:aws:iam::123456789012:role/example-role"

				  additional_tags {
					key   = "key1"
					value = "value1"
				  }

				  additional_tags {
					key   = "key2"
					value = "value2"
				  }

				  additional_tags {
					key   = "key3"
					value = "value3"
				  }
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "dynamodb_start_backup_job"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.dynamodb_table_name", "example-table"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.lifecycle_delete_after_days", "7"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.backup_vault_name", "example-vault"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.iam_role_arn", "arn:aws:iam::123456789012:role/example-role"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.additional_tags.0.key", "key1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.additional_tags.0.value", "value1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.additional_tags.1.key", "key2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.additional_tags.1.value", "value2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.additional_tags.2.key", "key3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "dynamodb_start_backup_job_action_value.0.additional_tags.2.value", "value3"),
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
