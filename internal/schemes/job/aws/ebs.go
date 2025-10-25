package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BulkDeleteEBSSnapshotsActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_by_tag_bulk_delete_ebs_snapshots": {
			Description: "Specifies whether to exclude EBS Snapshots with certain tags from deletion",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"exclude_by_tag_key_bulk_delete_ebs_snapshots": {
			Description: "The tag key used to identify EBS Snapshots to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"exclude_by_tag_value_bulk_delete_ebs_snapshots": {
			Description: "The tag value used to identify EBS Snapshots to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"specify_base_date": {
			Description: "Specifies the method for determining which EBS Snapshots to delete",
			Type:        schema.TypeString,
			Required:    true,
		},
		"before_days": {
			Description: "The number of days used to identify EBS Snapshots to be deleted",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"before_date": {
			Description: "The date used to identify EBS Snapshots for deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CopyEbsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "AWS Region from which the copy was made",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "AWS Region to copy to",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_ebs_snapshot": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_id": {
			Description: "Target EBS snapshot ID",
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
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func CreateEbsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_volume": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"volume_id": {
			Description: "対象のEBSボリュームID",
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
		"description": {
			Description: "Description to be set for EBS volume",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tags": {
			Description: "Array of tags to assign to the created EBS volume",
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
		"additional_tag_key": {
			Description: "Tag key to assign to the created EBS volume",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "Tag key to assign to the created EBS volume Tag value to assign to the created EBS volume",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
