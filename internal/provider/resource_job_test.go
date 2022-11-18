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

func TestAccCloudAutomatorJob_CronRuleOneTime(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCronRuleOneTime(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "cron"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.hour", "3"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.minutes", "30"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.schedule_type", "one_time"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.one_time_schedule", "2099/01/01"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.time_zone", "Tokyo"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CronRuleWeekly(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCronRuleWeekly(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "cron"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.hour", "3"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.minutes", "30"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.schedule_type", "weekly"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.weekly_schedule.#", "2"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.weekly_schedule.0", "monday"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.weekly_schedule.1", "sunday"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.time_zone", "Tokyo"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.dates_to_skip.0", "2099-12-31"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.national_holiday_schedule", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CronRuleMonthlyDayOfWeek(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCronRuleMonthlyDayOfWeek(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "cron"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.hour", "3"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.minutes", "30"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.schedule_type", "monthly_day_of_week"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.monthly_day_of_week_schedule.0.friday.0", "2"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.time_zone", "Tokyo"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.dates_to_skip.0", "2099-12-31"),
					resource.TestCheckResourceAttr(
						resourceName, "cron_rule_value.0.national_holiday_schedule", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_WebhookRule(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigWebhookRule(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "webhook"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_ScheduleRule(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigScheduleRule(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "schedule"),
					resource.TestCheckResourceAttr(
						resourceName, "schedule_rule_value.0.schedule", "2099-12-31 10:10:00\n2099-01-01 22:40:00"),
					resource.TestCheckResourceAttr(
						resourceName, "schedule_rule_value.0.time_zone", "Tokyo"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_SqsV2Rule(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	sqsAwsAccountId := acctest.TestSqsAwsAccountId()
	sqsRegion := acctest.TestSqsRegion()
	sqsQueue := acctest.TestSqsQueue()
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigSqsV2Rule(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "sqs_v2"),
					resource.TestCheckResourceAttr(
						resourceName, "sqs_v2_rule_value.0.sqs_aws_account_id", sqsAwsAccountId),
					resource.TestCheckResourceAttr(
						resourceName, "sqs_v2_rule_value.0.sqs_region", sqsRegion),
					resource.TestCheckResourceAttr(
						resourceName, "sqs_v2_rule_value.0.queue", sqsQueue),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_AmazonSnsRule(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigAmazonSnsRule(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "rule_type", "amazon_sns"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_AuthorizeSecurityGroupIngressAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigAuthorizeSecurityGroupIngressAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "authorize_security_group_ingress"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.specify_security_group", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.ip_protocol", "tcp"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.to_port", "80"),
					resource.TestCheckResourceAttr(
						resourceName, "authorize_security_group_ingress_action_value.0.cidr_ip", "172.31.0.0/16"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_ChangeRdsClusterInstanceClassAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigChangeRdsClusterInstanceClassAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "change_rds_cluster_instance_class"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_cluster_instance_class_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_cluster_instance_class_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_cluster_instance_class_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_cluster_instance_class_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_cluster_instance_class_action_value.0.db_instance_class", "db.t3.micro"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_ChangeRdsInstanceClassAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigChangeRdsInstanceClassAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "change_rds_instance_class"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_instance_class_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_instance_class_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_instance_class_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_instance_class_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "change_rds_instance_class_action_value.0.db_instance_class", "db.t3.micro"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_ChangeInstanceTypeAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigChangeInstanceTypeAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "change_instance_type"),
					resource.TestCheckResourceAttr(
						resourceName, "change_instance_type_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "change_instance_type_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "change_instance_type_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "change_instance_type_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "change_instance_type_action_value.0.instance_type", "t2.medium"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CopyEbsSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCopyEbsSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "copy_ebs_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_ebs_snapshot_action_value.0.source_region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_ebs_snapshot_action_value.0.destination_region", "us-east-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_ebs_snapshot_action_value.0.specify_ebs_snapshot", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_ebs_snapshot_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_ebs_snapshot_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_ebs_snapshot_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CopyImageAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCopyImageAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "copy_image"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_image_action_value.0.source_region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_image_action_value.0.destination_region", "us-east-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_image_action_value.0.specify_image", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_image_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_image_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_image_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CopyRdsClusterSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCopyRdsClusterSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "copy_rds_cluster_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_cluster_snapshot_action_value.0.source_region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_cluster_snapshot_action_value.0.destination_region", "ap-southeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_cluster_snapshot_action_value.0.specify_rds_cluster_snapshot", "rds_cluster_snapshot_id"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_cluster_snapshot_action_value.0.rds_cluster_snapshot_id", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_cluster_snapshot_action_value.0.kms_key_id", "1234abcd-12ab-34cd-56ef-1234567890ab"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CopyRdsSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCopyRdsSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "copy_rds_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_snapshot_action_value.0.source_region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_snapshot_action_value.0.destination_region", "us-east-1"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_snapshot_action_value.0.specify_rds_snapshot", "identifier"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_snapshot_action_value.0.rds_snapshot_id", "test-db"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_snapshot_action_value.0.option_group_name", "default:mysql-5-6"),
					resource.TestCheckResourceAttr(
						resourceName, "copy_rds_snapshot_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CreateFSxBackupAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCreateFSxBackupAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "create_fsx_backup"),
					resource.TestCheckResourceAttr(
						resourceName, "create_fsx_backup_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_fsx_backup_action_value.0.specify_file_system", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "create_fsx_backup_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "create_fsx_backup_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "create_fsx_backup_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "create_fsx_backup_action_value.0.backup_name", "example-backup"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CreateEbsSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCreateEbsSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "create_ebs_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.specify_volume", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.description", "test db"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.additional_tag_key", "example-key"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.additional_tag_value", "example-value"),
					resource.TestCheckResourceAttr(
						resourceName, "create_ebs_snapshot_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CreateImageAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCreateImageAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "create_image"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.specify_image_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.image_name", "test-image"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.description", "test image"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.reboot_instance", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.additional_tags.0.key", "key-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.additional_tags.0.value", "value-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.additional_tags.1.key", "key-2"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.additional_tags.1.value", "value-2"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.add_same_tag_to_snapshot", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "create_image_action_value.0.recreate_image_if_ami_status_failed", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CreateRdsClusterSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCreateRdsClusterSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "create_rds_cluster_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.specify_rds_cluster", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.db_cluster_snapshot_identifier", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_cluster_snapshot_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CreateRdsSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCreateRdsSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "create_rds_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.rds_snapshot_id", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "create_rds_snapshot_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_CreateRedshiftSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigCreateRedshiftSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "create_redshift_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.specify_cluster", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.cluster_snapshot_identifier", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "create_redshift_snapshot_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_DelayAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigDelayAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "delay"),
					resource.TestCheckResourceAttr(
						resourceName, "delay_action_value.0.delay_minutes", "30"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_DeleteClusterAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigDeleteClusterAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "delete_cluster"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_cluster_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_cluster_action_value.0.cluster_identifier", "test-cluster"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_cluster_action_value.0.final_cluster_snapshot_identifier", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_cluster_action_value.0.skip_final_cluster_snapshot", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_cluster_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_DeleteRdsClusterAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigDeleteRdsClusterAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "delete_rds_cluster"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.specify_rds_cluster", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.final_db_snapshot_identifier", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.skip_final_snapshot", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_cluster_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_DeleteRdsInstanceAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigDeleteRdsInstanceAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "delete_rds_instance"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.final_rds_snapshot_id", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.skip_final_rds_snapshot", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "delete_rds_instance_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_DeregisterInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigDeregisterInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "deregister_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_instances_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_instances_action_value.0.load_balancer_name", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_DeregisterTargetInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigDeregisterTargetInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "deregister_target_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_target_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_target_instances_action_value.0.target_group_arn", "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/t1/c8a1987f0402f55a"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_target_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "deregister_target_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_GoogleComputeInsertMachineImageAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigGoogleComputeInsertMachineImageAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "google_compute_insert_machine_image"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.region", "asia-northeast1"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.project_id", "example-project"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.specify_vm_instance", "label"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.vm_instance_label_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.vm_instance_label_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.machine_image_storage_location", "asia-northeast1"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.machine_image_basename", "example-daily"),
					resource.TestCheckResourceAttr(
						resourceName, "google_compute_insert_machine_image_action_value.0.generation", "10"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RebootRdsInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRebootRdsInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "reboot_rds_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_rds_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_rds_instances_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_rds_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_rds_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RebootWorkspacesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRebootWorkspacesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "reboot_workspaces"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_workspaces_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_workspaces_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "reboot_workspaces_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RebuildWorkspacesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRebuildWorkspacesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "rebuild_workspaces"),
					resource.TestCheckResourceAttr(
						resourceName, "rebuild_workspaces_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "rebuild_workspaces_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "rebuild_workspaces_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RegisterInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRegisterInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "register_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "register_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "register_instances_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "register_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "register_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "register_instances_action_value.0.load_balancer_name", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RegisterTargetInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRegisterTargetInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "register_target_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "register_target_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "register_target_instances_action_value.0.target_group_arn", "arn:aws:elasticloadbalancing:ap-northeast-1:123456789012:targetgroup/t1/c8a1987f0402f55a"),
					resource.TestCheckResourceAttr(
						resourceName, "register_target_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "register_target_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RestoreFromClusterSnapshotAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRestoreFromClusterSnapshotAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "restore_from_cluster_snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.cluster_identifier", "test-cluster"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.snapshot_identifier", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.cluster_parameter_group_name", "test-parameter-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.cluster_subnet_group_name", "test-subnet-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.port", "5432"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.publicly_accessible", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.availability_zone", "ap-northeast-1a"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.vpc_security_group_ids.0", "sg-00000001"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.vpc_security_group_ids.1", "sg-00000002"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.allow_version_upgrade", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_from_cluster_snapshot_action_value.0.delete_cluster_snapshot", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RestoreRdsClusterAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRestoreRdsClusterAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "restore_rds_cluster"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.db_instance_identifier", "test-db-instance"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.db_cluster_identifier", "test-cluster"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.snapshot_identifier", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.engine", "aurora"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.engine_version", "1.2.3.4"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.db_instance_class", "db.t2.micro"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.db_subnet_group_name", "test-subnet-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.publicly_accessible", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.availability_zone", "ap-northeast-1a"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.vpc_security_group_ids.0", "sg-00000001"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.vpc_security_group_ids.1", "sg-00000002"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.port", "5432"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.db_cluster_parameter_group_name", "test-cluster-parameter-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.db_parameter_group_name", "test-parameter-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.option_group_name", "test-option-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.auto_minor_version_upgrade", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_cluster_action_value.0.delete_db_cluster_snapshot", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RestoreRdsInstanceAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRestoreRdsInstanceAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "restore_rds_instance"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.rds_instance_id", "test-db-instance"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.rds_snapshot_id", "test-snapshot"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.db_engine", "mysql"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.license_model", "license-included"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.db_instance_class", "db.t2.micro"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.multi_az", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.storage_type", "gp2"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.iops", "30000"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.vpc", "vpc-00000001"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.subnet_group", "test-subnet-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.publicly_accessible", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.vpc_security_group_ids.0", "sg-00000001"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.vpc_security_group_ids.1", "sg-00000002"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.db_name", "testdb"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.port", "5432"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.parameter_group", "test-parameter-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.option_group", "test-option-group"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.auto_minor_version_upgrade", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.delete_rds_snapshot", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.additional_tag_key", "test-key"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.additional_tag_value", "test-value"),
					resource.TestCheckResourceAttr(
						resourceName, "restore_rds_instance_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_RevokeSecurityGroupIngressAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigRevokeSecurityGroupIngressAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "revoke_security_group_ingress"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.specify_security_group", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.ip_protocol", "tcp"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.to_port", "80"),
					resource.TestCheckResourceAttr(
						resourceName, "revoke_security_group_ingress_action_value.0.cidr_ip", "172.31.0.0/16"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_SendCommandAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigSendCommandAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "send_command"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.command", "whoami"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.comment", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.document_name", "AWS-RunShellScript"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.output_s3_bucket_name", "test-s3-bucket"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.output_s3_key_prefix", "test-key"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.timeout_seconds", "60"),
					resource.TestCheckResourceAttr(
						resourceName, "send_command_action_value.0.execution_timeout_seconds", "60"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StartInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStartInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "start_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "start_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "start_instances_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "start_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "start_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "start_instances_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "start_instances_action_value.0.status_checks_enable", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StartRdsClustersAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStartRdsClustersAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "start_rds_clusters"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_clusters_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_clusters_action_value.0.specify_rds_cluster", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_clusters_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_clusters_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_clusters_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StartRdsInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStartRdsInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "start_rds_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_instances_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "start_rds_instances_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StopInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStopInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "stop_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_instances_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_instances_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StopRdsClustersAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStopRdsClustersAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "stop_rds_clusters"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_clusters_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_clusters_action_value.0.specify_rds_cluster", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_clusters_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_clusters_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_clusters_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StopRdsInstancesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStopRdsInstancesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "stop_rds_instances"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_instances_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_instances_action_value.0.specify_rds_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_instances_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_instances_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "stop_rds_instances_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_StartWorkspacesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigStartWorkspacesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "start_workspaces"),
					resource.TestCheckResourceAttr(
						resourceName, "start_workspaces_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "start_workspaces_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "start_workspaces_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_TerminateWorkspacesAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigTerminateWorkspacesAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "terminate_workspaces"),
					resource.TestCheckResourceAttr(
						resourceName, "terminate_workspaces_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "terminate_workspaces_action_value.0.specify_workspace", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "terminate_workspaces_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "terminate_workspaces_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "terminate_workspaces_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_UpdateRecordSetAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigUpdateRecordSetAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "update_record_set"),
					resource.TestCheckResourceAttr(
						resourceName, "update_record_set_action_value.0.zone_name", "test.local."),
					resource.TestCheckResourceAttr(
						resourceName, "update_record_set_action_value.0.record_set_name", "aaa.test.local."),
					resource.TestCheckResourceAttr(
						resourceName, "update_record_set_action_value.0.record_set_type", "A"),
					resource.TestCheckResourceAttr(
						resourceName, "update_record_set_action_value.0.record_set_value", "1.2.3.4"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_WindowsUpdateAction(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigWindowsUpdateAction(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "windows_update"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.comment", "test"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.document_name", "AWS-InstallMissingWindowsUpdates"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.kb_article_ids", "KB1111111,KB2222222"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.output_s3_bucket_name", "test-s3-bucket"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.output_s3_key_prefix", "test-key"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.update_level", "All"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_action_value.0.timeout_seconds", "60"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_WindowsUpdateV2Action(t *testing.T) {
	resourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))
	postProcessId := acctest.TestPostProcessId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorJobConfigWindowsUpdateV2Action(jobName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						resourceName, "action_type", "windows_update_v2"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.specify_instance", "tag"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.tag_key", "env"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.tag_value", "develop"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.allow_reboot", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.severity_levels.0", "Critical"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.severity_levels.1", "Low"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.output_s3_bucket_name", "test-s3-bucket"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.output_s3_key_prefix", "test-key"),
					resource.TestCheckResourceAttr(
						resourceName, "windows_update_v2_action_value.0.trace_status", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "completed_post_process_id.0", postProcessId),
					resource.TestCheckResourceAttr(
						resourceName, "failed_post_process_id.0", postProcessId),
				),
			},
		},
	})
}

func testAccCheckCloudAutomatorJobExists(accProvider *schema.Provider, n string) resource.TestCheckFunc {
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

func testAccCheckCloudAutomatorJobConfigCronRuleOneTime(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCronRuleWeekly(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCronRuleMonthlyDayOfWeek(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigWebhookRule(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigScheduleRule(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigSqsV2Rule(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestSqsAwsAccountId(), acctest.TestSqsRegion(), acctest.TestSqsQueue(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigAmazonSnsRule(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigNoRule(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigAuthorizeSecurityGroupIngressAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigChangeRdsClusterInstanceClassAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigChangeRdsInstanceClassAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigChangeInstanceTypeAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCopyEbsSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCopyImageAction(rName string) string {
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
		trace_status = "true"
	}
	completed_post_process_id = [%s]
	failed_post_process_id = [%s]
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCopyRdsClusterSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCopyRdsSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCreateFSxBackupAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCreateEbsSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCreateImageAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCreateRdsClusterSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCreateRdsSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigCreateRedshiftSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigDelayAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigDeleteClusterAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigDeleteRdsClusterAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigDeleteRdsInstanceAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigDeregisterInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigDeregisterTargetInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigGoogleComputeInsertMachineImageAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRebootRdsInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRebootWorkspacesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRebuildWorkspacesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRegisterInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRegisterTargetInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRestoreFromClusterSnapshotAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRestoreRdsClusterAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRestoreRdsInstanceAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigRevokeSecurityGroupIngressAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigSendCommandAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStartInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStartRdsClustersAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStartRdsInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStopInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStopRdsClustersAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStopRdsInstancesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigStartWorkspacesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigTerminateWorkspacesAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigUpdateRecordSetAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigWindowsUpdateAction(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}

func testAccCheckCloudAutomatorJobConfigWindowsUpdateV2Action(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
}
