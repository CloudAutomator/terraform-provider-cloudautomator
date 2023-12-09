package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func S3StartBackupJobActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"bucket_name": {
			Description: "Name of the S3 bucket to store the backup",
			Type:        schema.TypeString,
			Required:    true,
		},
		"backup_vault_name": {
			Description: "Name of the backup vault to store the backup",
			Type:        schema.TypeString,
			Required:    true,
		},
		"lifecycle_delete_after_days": {
			Description: "Number of days to keep the backup",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"iam_role_arn": {
			Description: "ARN of the IAM role to use for the backup",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "Additional tags to be added to the backup",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key": {
						Type:     schema.TypeString,
						Required: true,
					},
					"value": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}
