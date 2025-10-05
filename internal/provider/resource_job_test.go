package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/client"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudAutomatorJob(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "CronRuleOneTime",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "cron"
				cron_rule_value {
					hour = "3"
					minutes = "30"
					schedule_type = "one_time"
					one_time_schedule = "2099/01/01"
					time_zone = "Tokyo"
				}

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "cron"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.hour", "3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.minutes", "30"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.schedule_type", "one_time"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.one_time_schedule", "2099/01/01"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.time_zone", "Tokyo"),
			},
		},
		{
			name:    "CronRuleWeekly",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "cron"
				cron_rule_value {
					hour = "3"
					minutes = "30"
					schedule_type = "weekly"
					weekly_schedule = [
						"monday",
						"sunday"
					]
					time_zone = "Tokyo"
					dates_to_skip = ["2099-12-31"]
					national_holiday_schedule = "true"
				}

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "cron"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.hour", "3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.minutes", "30"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.schedule_type", "weekly"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.weekly_schedule.#", "2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.weekly_schedule.0", "monday"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.weekly_schedule.1", "sunday"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.time_zone", "Tokyo"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.dates_to_skip.0", "2099-12-31"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.national_holiday_schedule", "true"),
			},
		},
		{
			name:    "CronRuleMonthlyDayOfWeek",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "cron"
				cron_rule_value {
					hour = "3"
					minutes = "30"
					schedule_type = "monthly_day_of_week"
					monthly_day_of_week_schedule {
					  friday = [2]
					}
					time_zone = "Tokyo"
					dates_to_skip = ["2099-12-31"]
					national_holiday_schedule = "true"
				}

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "cron"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.hour", "3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.minutes", "30"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.schedule_type", "monthly_day_of_week"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.monthly_day_of_week_schedule.0.friday.0", "2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.time_zone", "Tokyo"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.dates_to_skip.0", "2099-12-31"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "cron_rule_value.0.national_holiday_schedule", "true"),
			},
		},
		{
			name:    "WebhookRule",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "webhook"

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "webhook"),
			},
		},
		{
			name:    "ScheduleRule",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "schedule"
				schedule_rule_value {
					schedule = "2099-12-31 10:10:00\n2099-01-01 22:40:00"
					time_zone = "Tokyo"
				}

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "schedule"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "schedule_rule_value.0.schedule", "2099-12-31 10:10:00\n2099-01-01 22:40:00"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "schedule_rule_value.0.time_zone", "Tokyo"),
			},
		},
		{
			name:    "SqsV2Rule",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "sqs_v2"
				sqs_v2_rule_value {
					sqs_aws_account_id = "%s"
					sqs_region = "%s"
					queue = "%s"
				}

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestRegion(), acctest.TestSqsQueue(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "sqs_v2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "sqs_v2_rule_value.0.sqs_aws_account_id", acctest.TestAwsAccountId()),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "sqs_v2_rule_value.0.sqs_region", acctest.TestRegion()),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "sqs_v2_rule_value.0.queue", acctest.TestSqsQueue()),
			},
		},
		{
			name:    "AmazonSnsRule",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "amazon_sns"

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "amazon_sns"),
			},
		},
		{
			name:    "NoRule",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				for_workflow = true

				rule_type = "no_rule"

				action_type = "delay"
				delay_action_value {
					delay_minutes = 1
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rule_type", "no_rule"),
			},
		},
		{
			name:    "AttachUserPolicyAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "attach_user_policy"
				attach_user_policy_action_value {
					user_name = "example-user"
					policy_arn = "arn:aws:iam::123456789012:policy/example-policy"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "attach_user_policy"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "attach_user_policy_action_value.0.user_name", "example-user"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "attach_user_policy_action_value.0.policy_arn", "arn:aws:iam::123456789012:policy/example-policy"),
			},
		},
		{
			name:    "AuthorizeSecurityGroupIngressAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "authorize_security_group_ingress"
				authorize_security_group_ingress_action_value {
					region = "ap-northeast-1"
					specify_security_group = "tag"
					tag_key = "env"
					tag_value = "develop"
					ip_protocol = "tcp"
					to_port = "80"
					cidr_ip = "172.31.0.0/16"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "authorize_security_group_ingress"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.specify_security_group", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.ip_protocol", "tcp"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.to_port", "80"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.cidr_ip", "172.31.0.0/16"),
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
		{
			name:    "BulkDeleteImagesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_ids = [%s]

				rule_type = "webhook"

				action_type = "bulk_delete_images"
				bulk_delete_images_action_value {
					exclude_by_tag_bulk_delete_images = true
					exclude_by_tag_key_bulk_delete_images = "env"
					exclude_by_tag_value_bulk_delete_images = "production"
					specify_base_date = "before_days"
					before_days = 365
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "bulk_delete_images"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.exclude_by_tag_bulk_delete_images", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.exclude_by_tag_key_bulk_delete_images", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.exclude_by_tag_value_bulk_delete_images", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.specify_base_date", "before_days"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.before_days", "365"),
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
			name:    "CopyImageAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "copy_image"
				copy_image_action_value {
					source_region = "ap-northeast-1"
					destination_region = "us-east-1"
					specify_image = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "copy_image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.source_region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.destination_region", "us-east-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.specify_image", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.trace_status", "true"),
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
			name:    "CreateFSxBackupAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_fsx_backup"
				create_fsx_backup_action_value {
					region = "ap-northeast-1"
					specify_file_system = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					backup_name = "example-backup"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_fsx_backup"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_fsx_backup_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_fsx_backup_action_value.0.specify_file_system", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_fsx_backup_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_fsx_backup_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_fsx_backup_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_fsx_backup_action_value.0.backup_name", "example-backup"),
			},
		},
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
			name:    "CreateImageAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_image"
				create_image_action_value {
					region = "ap-northeast-1"
					specify_image_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					image_name = "test-image"
					description = "test image"
					reboot_instance = "true"
					additional_tags {
						key = "key-1"
						value = "value-1"
					}
					additional_tags {
						key = "key-2"
						value = "value-2"
					}
					add_same_tag_to_snapshot = "true"
					trace_status = "true"
					recreate_image_if_ami_status_failed	 = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.specify_image_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.image_name", "test-image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.description", "test image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.reboot_instance", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.0.key", "key-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.0.value", "value-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.1.key", "key-2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.1.value", "value-2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.add_same_tag_to_snapshot", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.trace_status", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.recreate_image_if_ami_status_failed", "true"),
			},
		},
		{
			name:    "CreateNatGatewayAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_nat_gateway"
				create_nat_gateway_action_value {
				  region           = "ap-northeast-1"
				  allocation_id    = "eipalloc-0123456789abcdef0"
				  nat_gateway_name = "test-nat-gateway"
				  subnet_id        = "subnet-0123456789abcdef0"
				  route_table_id   = "rtb-0123456789abcdef0"

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
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_nat_gateway"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.allocation_id", "eipalloc-0123456789abcdef0"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.nat_gateway_name", "test-nat-gateway"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.subnet_id", "subnet-0123456789abcdef0"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.route_table_id", "rtb-0123456789abcdef0"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.0.key", "key1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.0.value", "value1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.1.key", "key2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.1.value", "value2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.2.key", "key3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.2.value", "value3"),
			},
		},
		{
			name:    "DeleteNatGatewayAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "delete_nat_gateway"
				delete_nat_gateway_action_value {
				  region    = "ap-northeast-1"
				  tag_key   = "Name"
				  tag_value = "test-nat-gateway"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delete_nat_gateway"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_nat_gateway_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_nat_gateway_action_value.0.tag_key", "Name"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_nat_gateway_action_value.0.tag_value", "test-nat-gateway"),
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
			name:    "DelayAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "webhook"

				action_type = "delay"
				delay_action_value {
					delay_minutes = 30
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delay"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delay_action_value.0.delay_minutes", "30"),
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
			name:    "DeregisterInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "deregister_instances"
				deregister_instances_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					load_balancer_name = "test"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "deregister_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_instances_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_instances_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "deregister_instances_action_value.0.load_balancer_name", "test"),
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
		{
			name:    "DetachUserPolicyAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "detach_user_policy"
				detach_user_policy_action_value {
					user_name = "example-user"
					policy_arn = "arn:aws:iam::123456789012:policy/example-policy"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "detach_user_policy"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "detach_user_policy_action_value.0.user_name", "example-user"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "detach_user_policy_action_value.0.policy_arn", "arn:aws:iam::123456789012:policy/example-policy"),
			},
		},
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
			name:    "EfsStartBackupJobAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "efs_start_backup_job"
				efs_start_backup_job_action_value {
					region                      = "ap-northeast-1"
					file_system_id              = "fs-abcdefg1234567890"
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
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "efs_start_backup_job"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.file_system_id", "fs-abcdefg1234567890"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.lifecycle_delete_after_days", "7"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.backup_vault_name", "example-vault"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.iam_role_arn", "arn:aws:iam::123456789012:role/example-role"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.additional_tags.0.key", "key1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.additional_tags.0.value", "value1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.additional_tags.1.key", "key2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.additional_tags.1.value", "value2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.additional_tags.2.key", "key3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "efs_start_backup_job_action_value.0.additional_tags.2.value", "value3"),
			},
		},
		{
			name:    "GoogleComputeInsertMachineImageAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				google_cloud_account_id = "%s"

				rule_type = "webhook"

				action_type = "google_compute_insert_machine_image"
				google_compute_insert_machine_image_action_value {
					region = "asia-northeast1"
					project_id = "example-project"
					specify_vm_instance = "label"
					vm_instance_label_key = "env"
					vm_instance_label_value = "develop"
					machine_image_storage_location = "asia-northeast1"
					machine_image_basename = "example-daily"
					generation = 10
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "google_compute_insert_machine_image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.region", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.project_id", "example-project"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.specify_vm_instance", "label"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.vm_instance_label_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.vm_instance_label_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.machine_image_storage_location", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.machine_image_basename", "example-daily"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.generation", "10"),
			},
		},
		{
			name:    "GoogleComputeStartVmInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
			  name = "%s"
			  group_id = "%s"
			  google_cloud_account_id = "%s"

			  rule_type = "webhook"

			  action_type = "google_compute_start_vm_instances"
			  google_compute_start_vm_instances_action_value {
				region = "asia-northeast1"
				project_id = "example-project"
				specify_vm_instance = "label"
					vm_instance_label_key = "env"
				vm_instance_label_value = "develop"
			  }
			  completed_post_process_id = [%s]
			  failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "google_compute_start_vm_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.region", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.project_id", "example-project"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.specify_vm_instance", "label"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.vm_instance_label_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.vm_instance_label_value", "develop"),
			},
		},
		{
			name:    "GoogleComputeStopVmInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
			  name = "%s"
			  group_id = "%s"
			  google_cloud_account_id = "%s"

			  rule_type = "webhook"

			  action_type = "google_compute_stop_vm_instances"
			  google_compute_stop_vm_instances_action_value {
				region = "asia-northeast1"
				project_id = "example-project"
				specify_vm_instance = "label"
					vm_instance_label_key = "env"
				vm_instance_label_value = "develop"
			  }
			  completed_post_process_id = [%s]
			  failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "google_compute_stop_vm_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.region", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.project_id", "example-project"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.specify_vm_instance", "label"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.vm_instance_label_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.vm_instance_label_value", "develop"),
			},
		},
		{
			name:    "InvokeLambdaFunctionAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "invoke_lambda_function"
				invoke_lambda_function_action_value {
					region = "ap-northeast-1"
					function_name = "test-function"
					payload = "{}"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "invoke_lambda_function"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "invoke_lambda_function_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "invoke_lambda_function_action_value.0.function_name", "test-function"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "invoke_lambda_function_action_value.0.payload", "{}"),
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
			name:    "RebootWorkspacesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "reboot_workspaces"
				reboot_workspaces_action_value {
					region = "ap-northeast-1"
					tag_key = "env"
					tag_value = "develop"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "reboot_workspaces"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_workspaces_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_workspaces_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "reboot_workspaces_action_value.0.tag_value", "develop"),
			},
		},
		{
			name:    "RebuildWorkspacesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "rebuild_workspaces"
				rebuild_workspaces_action_value {
					region = "ap-northeast-1"
					tag_key = "env"
					tag_value = "develop"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "rebuild_workspaces"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rebuild_workspaces_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rebuild_workspaces_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "rebuild_workspaces_action_value.0.tag_value", "develop"),
			},
		},
		{
			name:    "RegisterInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "register_instances"
				register_instances_action_value {
					region = "ap-northeast-1"
					specify_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					load_balancer_name = "test"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "register_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_instances_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_instances_action_value.0.specify_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_instances_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_instances_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "register_instances_action_value.0.load_balancer_name", "test"),
			},
		},
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
			name:    "RevokeSecurityGroupIngressAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "revoke_security_group_ingress"
				revoke_security_group_ingress_action_value {
					region = "ap-northeast-1"
					specify_security_group = "tag"
					tag_key = "env"
					tag_value = "develop"
					ip_protocol = "tcp"
					to_port = "80"
					cidr_ip = "172.31.0.0/16"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "revoke_security_group_ingress"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.specify_security_group", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.ip_protocol", "tcp"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.to_port", "80"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.cidr_ip", "172.31.0.0/16"),
			},
		},
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
				testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], "cloudautomator_job.test"),
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
			name:    "StartWorkspacesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "start_workspaces"
				start_workspaces_action_value {
					region = "ap-northeast-1"
					tag_key = "env"
					tag_value = "develop"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "start_workspaces"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_workspaces_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_workspaces_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "start_workspaces_action_value.0.tag_value", "develop"),
			},
		},
		{
			name:    "TerminateWorkspacesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "terminate_workspaces"
				terminate_workspaces_action_value {
					region = "ap-northeast-1"
					specify_workspace = "tag"
					tag_key = "env"
					tag_value = "develop"
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "terminate_workspaces"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "terminate_workspaces_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "terminate_workspaces_action_value.0.specify_workspace", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "terminate_workspaces_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "terminate_workspaces_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "terminate_workspaces_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "UpdateRecordSetAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "update_record_set"
				update_record_set_action_value {
					zone_name = "test.local."
					record_set_name = "aaa.test.local."
					record_set_type = "A"
					record_set_value = "1.2.3.4"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "update_record_set"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "update_record_set_action_value.0.zone_name", "test.local."),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "update_record_set_action_value.0.record_set_name", "aaa.test.local."),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "update_record_set_action_value.0.record_set_type", "A"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "update_record_set_action_value.0.record_set_value", "1.2.3.4"),
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
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(tc.jobName),
						Check: resource.ComposeTestCheckFunc(
							append([]resource.TestCheckFunc{
								testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], "cloudautomator_job.test"),
								resource.TestCheckResourceAttr("cloudautomator_job.test", "name", tc.jobName),
								resource.TestCheckResourceAttr("cloudautomator_job.test", "completed_post_process_id.0", acctest.TestPostProcessId()),
								resource.TestCheckResourceAttr("cloudautomator_job.test", "failed_post_process_id.0", acctest.TestPostProcessId()),
							}, tc.checks...)...,
						),
					},
				},
			})
		})
	}
}

func testAccCheckCloudAutomatorJobExists(_ *schema.Provider, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*client.Client)

		if err := cloudautomatorJobExistsHelper(s, c, n); err != nil {
			return err
		}

		return nil
	}
}

func cloudautomatorJobExistsHelper(s *terraform.State, c *client.Client, name string) error {
	id := s.RootModule().Resources[name].Primary.ID
	if _, _, err := c.GetJob(id); err != nil {
		return fmt.Errorf("received an error retrieving job %s", err)
	}

	return nil
}

func testAccCheckCloudAutomatorJobDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudautomator_job" {
			continue
		}

		jobId := rs.Primary.ID

		_, res, err := c.GetJob(jobId)
		if err != nil {
			if res.StatusCode == 404 {
				continue
			}

			return fmt.Errorf("received an error retrieving job %s", err)
		}

		return fmt.Errorf("job exists.")
	}

	return nil
}
