package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EmailParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_recipient": {
			Description: "Email Address",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func SlackParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"slack_channel_name": {
			Description: "Slack Channel Name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"slack_language": {
			Description: "Language of notification content",
			Type:        schema.TypeString,
			Required:    true,
		},
		"slack_time_zone": {
			Description: "Time zone",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func SqsParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sqs_aws_account_id": {
			Description: "ID of the AWS account used to search the SQS queue",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"sqs_queue": {
			Description: "SQS queue name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"sqs_region": {
			Description: "Region name where the SQS queue resides",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func WebhookParametersFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"webhook_authorization_header": {
			Description: "Authorization header value",
			Type:        schema.TypeString,
			Required:    true,
		},
		"webhook_url": {
			Description: "URL of the Webhook destination",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
