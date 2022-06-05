package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CronRuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hour": {
			Description: "ジョブを実行するタイミング(時)",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"minutes": {
			Description: "ジョブを実行するタイミング(分)",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"schedule_type": {
			Description: "スケジュールのタイプ (one_time(一度きり) / weekly(毎週) / monthly(毎月の日付) / monthly_day_of_week(毎月の曜日))",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"one_time_schedule": {
			Description: "ジョブの実行年月日 (yyyy/mm/dd形式) ※`schedule_type` が `one_time` の場合必須",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"weekly_schedule": {
			Description: "ジョブの実行曜日 (sunday〜saturdayまでの文字列を含む配列) ※`schedule_type` が `weekly` の場合必須",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"monthly_schedule": {
			Description: "毎月のジョブの実行日付 (1〜31の日付または月末をあらわす `end_of_month` を文字列で指定する) ※`schedule_type` が `monthly` の場合必須",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"monthly_day_of_week_schedule": {
			Description: "毎月のジョブの曜日 ({ friday: [2] } のように曜日名をキー、順序の配列を値として指定する) ※`schedule_type` が `monthly_day_of_week` の場合必須",
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
			Description: "ジョブの実行日と日本の祝日が重なっていた場合にジョブの実行をスキップするか否かのフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"time_zone": {
			Description: "`Tokyo`, `Singapore`, `UTC` など、タイムゾーンを表す文字列",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"dates_to_skip": {
			Description: "ジョブ実行をスキップする日付を示す YYYY-MM-DD 形式の日付の配列",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"start_timeout_minutes": {
			Description: "ジョブの開始が遅延した場合にジョブ実行の開始をキャンセルする遅延時間",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func ScheduleRuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"schedule": {
			Description: "ジョブの実行予定日時 (YYYY-MM-DD HH:MM:SS の形式で `\n` 区切りの文字列)",
			Type:        schema.TypeString,
			Required:    true,
		},
		"time_zone": {
			Description: "実行予定日時のタイムゾーン (指定しない場合はジョブ作成ユーザーのタイムゾーンが設定されます)",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func SqsV2RuleValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sqs_aws_account_id": {
			Description: "SQSキューが存在するAWSアカウントのID (AWSのアカウントIDではなくCloud Automator上のID)",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"sqs_region": {
			Description: "SQSキューが存在するリージョン <br>例) “ap-northeast-1”",
			Type:        schema.TypeString,
			Required:    true,
		},
		"queue": {
			Description: "SQSキュー名 (標準キューのみ対応)",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
