package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EmailParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_recipient": {
			Description: "メールアドレス",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func SlackParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"slack_channel_name": {
			Description: "Slackチャンネル名",
			Type:        schema.TypeString,
			Required:    true,
		},
		"slack_language": {
			Description: "通知内容の言語",
			Type:        schema.TypeString,
			Required:    true,
		},
		"slack_time_zone": {
			Description: "通知内容のタイムゾーン",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func SqsParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sqs_aws_account_id": {
			Description: "SQSのキューを検索する際に利用するAWSアカウントのID",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"sqs_queue": {
			Description: "SQSのキュー名",
			Type:        schema.TypeString,
			Required:    true,
		},
		"sqs_region": {
			Description: "SQSのキューが存在するリージョン名",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func WebhookParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"webhook_authorization_header": {
			Description: "Authorizationヘッダの値",
			Type:        schema.TypeString,
			Required:    true,
		},
		"webhook_url": {
			Description: "Webhook送信先となるURL",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
