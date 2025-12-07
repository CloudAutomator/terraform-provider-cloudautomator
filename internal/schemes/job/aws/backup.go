package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func VaultRecoveryPointStartCopyJobActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "Region where the source backup vault is located",
			Type:        schema.TypeString,
			Required:    true,
		},
		"source_backup_vault_name": {
			Description: "Source backup vault name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"resource_type": {
			Description: "Type of resource to copy recovery point",
			Type:        schema.TypeString,
			Required:    true,
		},
		"resource_id": {
			Description: "ID of resource to copy recovery point",
			Type:        schema.TypeString,
			Required:    true,
		},
		"iam_role_arn": {
			Description: "IAM Role ARN to use when creating the copy",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_destination_aws_account": {
			Description: "How to specify the destination AWS account",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "Region where the destination backup vault is located",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"destination_backup_vault_name": {
			Description: "Destination backup vault name",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"destination_backup_vault_arn": {
			Description: "ARN of the destination backup vault",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"lifecycle_delete_after_days": {
			Description: "Copy retention period (days)",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}
