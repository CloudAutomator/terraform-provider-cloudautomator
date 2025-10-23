package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Ebs(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "CreateEbsSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_ebs_snapshot"
				create_ebs_snapshot_action_value {
					region = "ap-northeast-1"
					specify_volume = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					description = "test db"
					additional_tag_key = "example-key"
					additional_tag_value = "example-value"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_ebs_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.specify_volume", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.description", "test db"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.additional_tag_key", "example-key"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.additional_tag_value", "example-value"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_ebs_snapshot_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "CopyEbsSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "copy_ebs_snapshot"
				copy_ebs_snapshot_action_value {
					source_region = "ap-northeast-1"
					destination_region = "us-east-1"
					specify_ebs_snapshot = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "copy_ebs_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_ebs_snapshot_action_value.0.source_region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_ebs_snapshot_action_value.0.destination_region", "us-east-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_ebs_snapshot_action_value.0.specify_ebs_snapshot", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_ebs_snapshot_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_ebs_snapshot_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_ebs_snapshot_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "BulkDeleteEBSSnapshotsAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_ids = [%s]

				rule_type = "webhook"

				action_type = "bulk_delete_ebs_snapshots"
				bulk_delete_ebs_snapshots_action_value {
					exclude_by_tag_bulk_delete_ebs_snapshots = true
					exclude_by_tag_key_bulk_delete_ebs_snapshots = "env"
					exclude_by_tag_value_bulk_delete_ebs_snapshots = "production"
					specify_base_date = "before_days"
					before_days = 365
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "bulk_delete_ebs_snapshots"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_ebs_snapshots_action_value.0.exclude_by_tag_bulk_delete_ebs_snapshots", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_ebs_snapshots_action_value.0.exclude_by_tag_key_bulk_delete_ebs_snapshots", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_ebs_snapshots_action_value.0.exclude_by_tag_value_bulk_delete_ebs_snapshots", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_ebs_snapshots_action_value.0.specify_base_date", "before_days"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_ebs_snapshots_action_value.0.before_days", "365"),
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
