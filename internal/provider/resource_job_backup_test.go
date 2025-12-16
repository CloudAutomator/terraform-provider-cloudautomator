package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Backup(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "VaultRecoveryPointStartCopyJobAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "vault_recovery_point_start_copy_job"
				vault_recovery_point_start_copy_job_action_value {
				  source_region                       = "ap-northeast-1"
				  source_backup_vault_name            = "source-vault"
				  resource_type                       = "EC2"
				  resource_id                         = "i-00000001"
				  iam_role_arn                        = "arn:aws:iam::123456789012:role/RoleForCopy"
				  specify_destination_aws_account     = "same"
				  destination_region                  = "ap-southeast-1"
				  destination_backup_vault_name       = "dest-vault"
				  lifecycle_delete_after_days         = 7
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "vault_recovery_point_start_copy_job"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.source_region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.source_backup_vault_name", "source-vault"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.resource_type", "EC2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.resource_id", "i-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.iam_role_arn", "arn:aws:iam::123456789012:role/RoleForCopy"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.specify_destination_aws_account", "same"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.destination_region", "ap-southeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.destination_backup_vault_name", "dest-vault"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "vault_recovery_point_start_copy_job_action_value.0.lifecycle_delete_after_days", "7"),
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
