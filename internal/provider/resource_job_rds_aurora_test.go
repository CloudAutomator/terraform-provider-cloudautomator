package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_RdsAurora(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "StartRdsClustersAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "start_rds_clusters"
				start_rds_clusters_action_value {
					region = "ap-northeast-1"
					specify_rds_cluster = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "start_rds_clusters"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_clusters_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_clusters_action_value.0.specify_rds_cluster", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_clusters_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_clusters_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_clusters_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "StopRdsClustersAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "stop_rds_clusters"
				stop_rds_clusters_action_value {
					region = "ap-northeast-1"
					specify_rds_cluster = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "stop_rds_clusters"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_clusters_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_clusters_action_value.0.specify_rds_cluster", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_clusters_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_clusters_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_clusters_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "DeleteRdsClusterAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "delete_rds_cluster"
				delete_rds_cluster_action_value {
					region = "ap-northeast-1"
					specify_rds_cluster = "tag"
					tag_key = "env"
					tag_value = "develop"
					final_db_snapshot_identifier = "test-snapshot"
					skip_final_snapshot = "false"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delete_rds_cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.specify_rds_cluster", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.final_db_snapshot_identifier", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.skip_final_snapshot", "false"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_cluster_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "CreateRdsClusterSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_rds_cluster_snapshot"
				create_rds_cluster_snapshot_action_value {
					region = "ap-northeast-1"
					specify_rds_cluster = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					db_cluster_snapshot_identifier = "test"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_rds_cluster_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.specify_rds_cluster", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.db_cluster_snapshot_identifier", "test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_cluster_snapshot_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "CopyRdsClusterSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "copy_rds_cluster_snapshot"
				copy_rds_cluster_snapshot_action_value {
				  source_region = "ap-northeast-1"
				  destination_region = "ap-southeast-1"
				  specify_rds_cluster_snapshot = "rds_cluster_snapshot_id"
				  rds_cluster_snapshot_id = "test-snapshot"
				  kms_key_id = "1234abcd-12ab-34cd-56ef-1234567890ab"
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "copy_rds_cluster_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_cluster_snapshot_action_value.0.source_region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_cluster_snapshot_action_value.0.destination_region", "ap-southeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_cluster_snapshot_action_value.0.specify_rds_cluster_snapshot", "rds_cluster_snapshot_id"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_cluster_snapshot_action_value.0.rds_cluster_snapshot_id", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_cluster_snapshot_action_value.0.kms_key_id", "1234abcd-12ab-34cd-56ef-1234567890ab"),
			},
		},
		{
			name:    "RestoreRdsClusterAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "restore_rds_cluster"
				restore_rds_cluster_action_value {
					region = "ap-northeast-1"
					db_instance_identifier = "test-db-instance"
					db_cluster_identifier = "test-cluster"
					snapshot_identifier = "test-snapshot"
					engine = "aurora"
					engine_version = "1.2.3.4"
					db_instance_class = "db.t2.micro"
					db_subnet_group_name = "test-subnet-group"
					publicly_accessible = "true"
					availability_zone = "ap-northeast-1a"
					vpc_security_group_ids = [
						"sg-00000001",
						"sg-00000002"
					]
					port = 5432
					db_cluster_parameter_group_name = "test-cluster-parameter-group"
					db_parameter_group_name = "test-parameter-group"
					option_group_name = "test-option-group"
					auto_minor_version_upgrade = "true"
					delete_db_cluster_snapshot = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "restore_rds_cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.db_instance_identifier", "test-db-instance"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.db_cluster_identifier", "test-cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.snapshot_identifier", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.engine", "aurora"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.engine_version", "1.2.3.4"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.db_instance_class", "db.t2.micro"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.db_subnet_group_name", "test-subnet-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.publicly_accessible", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.availability_zone", "ap-northeast-1a"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.vpc_security_group_ids.0", "sg-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.vpc_security_group_ids.1", "sg-00000002"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.port", "5432"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.db_cluster_parameter_group_name", "test-cluster-parameter-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.db_parameter_group_name", "test-parameter-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.option_group_name", "test-option-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.auto_minor_version_upgrade", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_cluster_action_value.0.delete_db_cluster_snapshot", "true"),
			},
		},
		{
			name:    "RestoreFromClusterSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "restore_from_cluster_snapshot"
				restore_from_cluster_snapshot_action_value {
					region = "ap-northeast-1"
					cluster_identifier = "test-cluster"
					snapshot_identifier = "test-snapshot"
					cluster_parameter_group_name = "test-parameter-group"
					cluster_subnet_group_name = "test-subnet-group"
					port = 5432
					publicly_accessible = "true"
					availability_zone = "ap-northeast-1a"
					vpc_security_group_ids = [
						"sg-00000001",
						"sg-00000002"
					]
					allow_version_upgrade = "true"
					delete_cluster_snapshot = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "restore_from_cluster_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.cluster_identifier", "test-cluster"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.snapshot_identifier", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.cluster_parameter_group_name", "test-parameter-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.cluster_subnet_group_name", "test-subnet-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.port", "5432"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.publicly_accessible", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.availability_zone", "ap-northeast-1a"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.vpc_security_group_ids.0", "sg-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.vpc_security_group_ids.1", "sg-00000002"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.allow_version_upgrade", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_from_cluster_snapshot_action_value.0.delete_cluster_snapshot", "true"),
			},
		},
		{
			name:    "ChangeRdsClusterInstanceClassAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "change_rds_cluster_instance_class"
				change_rds_cluster_instance_class_action_value {
					region = "ap-northeast-1"
					specify_rds_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					db_instance_class = "db.t3.micro"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "change_rds_cluster_instance_class"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_cluster_instance_class_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_cluster_instance_class_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_cluster_instance_class_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_cluster_instance_class_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_cluster_instance_class_action_value.0.db_instance_class", "db.t3.micro"),
			},
		},
		{
			name:    "BulkDeleteRdsClusterSnapshotsAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_ids = [%s]

				rule_type = "webhook"

				action_type = "bulk_delete_rds_cluster_snapshots"
				bulk_delete_rds_cluster_snapshots_action_value {
					exclude_by_tag_bulk_delete_rds_cluster_snapshots = true
					exclude_by_tag_key_bulk_delete_rds_cluster_snapshots = "env"
					exclude_by_tag_value_bulk_delete_rds_cluster_snapshots = "production"
					specify_base_date = "before_days"
					before_days = 365
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "bulk_delete_rds_cluster_snapshots"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_cluster_snapshots_action_value.0.exclude_by_tag_bulk_delete_rds_cluster_snapshots", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_cluster_snapshots_action_value.0.exclude_by_tag_key_bulk_delete_rds_cluster_snapshots", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_cluster_snapshots_action_value.0.exclude_by_tag_value_bulk_delete_rds_cluster_snapshots", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_cluster_snapshots_action_value.0.specify_base_date", "before_days"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_cluster_snapshots_action_value.0.before_days", "365"),
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
