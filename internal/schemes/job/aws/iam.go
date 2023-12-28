package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AttachUserPolicyActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_name": {
			Description: "Name of the user to attach the policy to",
			Type:        schema.TypeString,
			Required:    true,
		},
		"policy_arn": {
			Description: "ARN of the policy to attach",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
