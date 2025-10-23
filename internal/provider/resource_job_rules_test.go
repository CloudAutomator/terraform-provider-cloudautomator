package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Rules(t *testing.T) {
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
