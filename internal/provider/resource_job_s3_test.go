package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_S3(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "S3StartBackupJobAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "s3_start_backup_job"
				s3_start_backup_job_action_value {
					region = "%s"
					bucket_name = "%s"
					backup_vault_name = "Default"
					lifecycle_delete_after_days = 7
					iam_role_arn = "arn:aws:iam::%s:role/service-role/AWSBackupDefaultServiceRole"
					additional_tags {
						key = "key-1"
						value= "value-1"
					}
					additional_tags {
						key = "key-2"
						value= "value-2"
					}
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestRegion(), acctest.TestS3BucketName(), acctest.TestAwsAccountNumber(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "s3_start_backup_job"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.region", acctest.TestRegion()),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.bucket_name", acctest.TestS3BucketName()),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.backup_vault_name", "Default"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.lifecycle_delete_after_days", "7"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.iam_role_arn", fmt.Sprintf("arn:aws:iam::%s:role/service-role/AWSBackupDefaultServiceRole", acctest.TestAwsAccountNumber())),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.additional_tags.0.key", "key-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.additional_tags.0.value", "value-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.additional_tags.1.key", "key-2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "s3_start_backup_job_action_value.0.additional_tags.1.value", "value-2"),
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
