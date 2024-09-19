package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CronRuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hour": {
			Description: "When to execute the job (hour)",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"minutes": {
			Description: "When to execute the job (minutes)",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"schedule_type": {
			Description: "Schedule type (one_time / weekly / monthly / monthly_day_of_week)",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"one_time_schedule": {
			Description: "Job execution date (yyyy/mm/dd format) *required if `schedule_type` is `one_time",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"weekly_schedule": {
			Description: "day of the week to run the job (array containing strings from sunday to saturday) *required if `schedule_type` is `weekly",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"monthly_schedule": {
			Description: "date to run the job each month (date between 1 and 31 or `end_of_month` as a string representing the end of the month) *required if `schedule_type` is `monthly`.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"monthly_day_of_week_schedule": {
			Description: "day of the week for monthly jobs (key is the name of the day and value is an ordered array, like { friday: [2] }) *required if `schedule_type` is `monthly_day_of_week`.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"monday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
					"tuesday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
					"wednesday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
					"thursday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
					"friday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
					"saturday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
					"sunday": {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
				},
			},
		},
		"national_holiday_schedule": {
			Description: "Whether to skip job execution if the job execution date coincides with a Japanese national holiday",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"time_zone": {
			Description: "A string representing the time zone, such as `Tokyo`, `Singapore`, `UTC`, etc.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dates_to_skip": {
			Description: "An array of dates in YYYYY-MM-DD format indicating the dates to skip job execution",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"start_timeout_minutes": {
			Description: "Delay time to cancel the start of job execution if the start of the job is delayed",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func ScheduleRuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"schedule": {
			Description: "The date and time the job is scheduled to run (a `\n` delimited string in the format YYYYY-MM-DD HH:MM:SS).",
			Type:        schema.TypeString,
			Required:    true,
		},
		"time_zone": {
			Description: "Time zone of the scheduled execution date and time (if not specified, the time zone of the user who created the job will be set)",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func AmazonSnsRuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"amazon_sns_token": {
			Description: "Amazon SNS token",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"time_zone": {
			Description: "Time zone",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func SqsV2RuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sqs_aws_account_id": {
			Description: "ID of the AWS account where the SQS queue resides (ID on Cloud Automator, not AWS account ID)",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"sqs_region": {
			Description: "Region where SQS queue exists <br> e.g.) \"ap-northeast-1\"",
			Type:        schema.TypeString,
			Required:    true,
		},
		"queue": {
			Description: "SQS queue name (only standard queues are supported)",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func WebhookRuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"token": {
			Description: "Token",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"time_zone": {
			Description: "Timezone",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"access_token": {
			Description: "Access token",
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
