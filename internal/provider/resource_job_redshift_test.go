package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Redshift(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "CreateRedshiftSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_redshift_snapshot"
				create_redshift_snapshot_action_value {
					region = "ap-northeast-1"
					specify_cluster = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					cluster_snapshot_identifier = "test"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_redshift_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.specify_cluster", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.cluster_snapshot_identifier", "test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_redshift_snapshot_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "DeleteClusterAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "delete_cluster"
				delete_cluster_action_value {
					region = "ap-northeast-1"
					cluster_identifier = "test-cluster"
					final_cluster_snapshot_identifier = "test-snapshot"
					skip_final_cluster_snapshot = "false"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delete_cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_cluster_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_cluster_action_value.0.cluster_identifier", "test-cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_cluster_action_value.0.final_cluster_snapshot_identifier", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_cluster_action_value.0.skip_final_cluster_snapshot", "false"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_cluster_action_value.0.trace_status", "true"),
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
