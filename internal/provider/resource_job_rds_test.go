package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Rds(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "StartRdsInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "start_rds_instances"
				start_rds_instances_action_value {
					region = "ap-northeast-1"
					specify_rds_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "start_rds_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_instances_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_instances_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_rds_instances_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "StopRdsInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "stop_rds_instances"
				stop_rds_instances_action_value {
					region = "ap-northeast-1"
					specify_rds_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "stop_rds_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_instances_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_instances_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "stop_rds_instances_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "RebootRdsInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "reboot_rds_instances"
				reboot_rds_instances_action_value {
					region = "ap-northeast-1"
					specify_rds_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "reboot_rds_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_rds_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_rds_instances_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_rds_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_rds_instances_action_value.0.tag_value", "develop"),
			},
		},
		{
			name:    "DeleteRdsInstanceAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "delete_rds_instance"
				delete_rds_instance_action_value {
					region = "ap-northeast-1"
					specify_rds_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					final_rds_snapshot_id = "test-snapshot"
					skip_final_rds_snapshot = "false"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delete_rds_instance"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.final_rds_snapshot_id", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.skip_final_rds_snapshot", "false"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_rds_instance_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "CreateRdsSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_rds_snapshot"
				create_rds_snapshot_action_value {
					region = "ap-northeast-1"
					specify_rds_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					rds_snapshot_id = "test"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_rds_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.rds_snapshot_id", "test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_rds_snapshot_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "CopyRdsSnapshotAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "copy_rds_snapshot"
				copy_rds_snapshot_action_value {
					source_region = "ap-northeast-1"
					destination_region = "us-east-1"
					specify_rds_snapshot = "identifier"
					rds_snapshot_id = "test-db"
					option_group_name = "default:mysql-5-6"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "copy_rds_snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_snapshot_action_value.0.source_region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_snapshot_action_value.0.destination_region", "us-east-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_snapshot_action_value.0.specify_rds_snapshot", "identifier"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_snapshot_action_value.0.rds_snapshot_id", "test-db"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_snapshot_action_value.0.option_group_name", "default:mysql-5-6"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_rds_snapshot_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "RestoreRdsInstanceAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "restore_rds_instance"
				restore_rds_instance_action_value {
					region = "ap-northeast-1"
					rds_instance_id = "test-db-instance"
					rds_snapshot_id = "test-snapshot"
					db_engine = "mysql"
					license_model = "license-included"
					db_instance_class = "db.t2.micro"
					multi_az = "true"
					storage_type = "gp2"
					iops = 30000
					vpc = "vpc-00000001"
					subnet_group = "test-subnet-group"
					publicly_accessible = "true"
					vpc_security_group_ids = [
						"sg-00000001",
						"sg-00000002"
					]
					db_name = "testdb"
					port = 5432
					parameter_group = "test-parameter-group"
					option_group = "test-option-group"
					auto_minor_version_upgrade = "true"
					delete_rds_snapshot = "true"
					additional_tag_key = "test-key"
					additional_tag_value = "test-value"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "restore_rds_instance"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.rds_instance_id", "test-db-instance"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.rds_snapshot_id", "test-snapshot"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.db_engine", "mysql"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.license_model", "license-included"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.db_instance_class", "db.t2.micro"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.multi_az", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.storage_type", "gp2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.iops", "30000"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.vpc", "vpc-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.subnet_group", "test-subnet-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.publicly_accessible", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.vpc_security_group_ids.0", "sg-00000001"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.vpc_security_group_ids.1", "sg-00000002"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.db_name", "testdb"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.port", "5432"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.parameter_group", "test-parameter-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.option_group", "test-option-group"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.auto_minor_version_upgrade", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.delete_rds_snapshot", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.additional_tag_key", "test-key"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.additional_tag_value", "test-value"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "restore_rds_instance_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "ChangeRdsInstanceClassAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "change_rds_instance_class"
				change_rds_instance_class_action_value {
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
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "change_rds_instance_class"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_instance_class_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_instance_class_action_value.0.specify_rds_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_instance_class_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_instance_class_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "change_rds_instance_class_action_value.0.db_instance_class", "db.t3.micro"),
			},
		},
		{
			name:    "Ec2StartBackupJobAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "ec2_start_backup_job"
				ec2_start_backup_job_action_value {
				  region                      = "ap-northeast-1"
				  specify_instance            = "tag"
          tag_key                     = "env"
          tag_value                   = "production"
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
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "ec2_start_backup_job"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.tag_value", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.lifecycle_delete_after_days", "7"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.backup_vault_name", "example-vault"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.iam_role_arn", "arn:aws:iam::123456789012:role/example-role"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.additional_tags.0.key", "key1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.additional_tags.0.value", "value1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.additional_tags.1.key", "key2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.additional_tags.1.value", "value2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.additional_tags.2.key", "key3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "ec2_start_backup_job_action_value.0.additional_tags.2.value", "value3"),
			},
		},
		{
			name:    "BulkDeleteRdsSnapshotsAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_ids = [%s]

				rule_type = "webhook"

				action_type = "bulk_delete_rds_snapshots"
				bulk_delete_rds_snapshots_action_value {
					exclude_by_tag_bulk_delete_rds_snapshots = true
					exclude_by_tag_key_bulk_delete_rds_snapshots = "env"
					exclude_by_tag_value_bulk_delete_rds_snapshots = "production"
					specify_base_date = "before_days"
					before_days = 365
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "bulk_delete_rds_snapshots"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_snapshots_action_value.0.exclude_by_tag_bulk_delete_rds_snapshots", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_snapshots_action_value.0.exclude_by_tag_key_bulk_delete_rds_snapshots", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_snapshots_action_value.0.exclude_by_tag_value_bulk_delete_rds_snapshots", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_snapshots_action_value.0.specify_base_date", "before_days"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_rds_snapshots_action_value.0.before_days", "365"),
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
