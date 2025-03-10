package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func EfsStartBackupJobActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region",
			Type:        schema.TypeString,
			Required:    true,
		},
		"file_system_id": {
			Description: "Target file system ID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"backup_vault_name": {
			Description: "Backup Vault Name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"lifecycle_delete_after_days": {
			Description: "Number of days to hold recovery point",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"iam_role_arn": {
			Description: "IAM Role ARN",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "Array of tags to be added to the recovery point",
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
