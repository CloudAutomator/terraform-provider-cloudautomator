package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CreateFSxBackupActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_file_system": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"file_system_id": {
			Description: "Target Filesystem snapshot ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "Tag key used to identify the target resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "Tag value used to identify the target resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"generation": {
			Description: "Number of EBS volumes to manage generation",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"backup_name": {
			Description: "Backup name",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
