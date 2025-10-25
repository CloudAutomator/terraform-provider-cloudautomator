package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BulkDeleteImagesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_by_tag_bulk_delete_images": {
			Description: "Specifies whether to exclude AMIs with certain tags from deletion",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"exclude_by_tag_key_bulk_delete_images": {
			Description: "The tag key used to identify AMIs to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"exclude_by_tag_value_bulk_delete_images": {
			Description: "The tag value used to identify AMIs to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"specify_base_date": {
			Description: "Specifies the method for determining which AMIs to delete",
			Type:        schema.TypeString,
			Required:    true,
		},
		"before_days": {
			Description: "The number of days used to identify AMIs to be deleted",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"before_date": {
			Description: "The date used to identify AMIs for deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CopyImageActionValueFields() map[string]*schema.Schema {
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
		"specify_image": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"source_image_id": {
			Description: "対象のAMIのID",
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
			Description: "Number of AMI generations to be managed (0-10). Required when specify_image is 'tag'",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateImageActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_image_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "Target EC2 instance ID",
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
			Description: "Number of AMI generations to be managed",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"image_name": {
			Description: "Image name to set for AMI",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description to be set in AMI",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"reboot_instance": {
			Description: "Flag whether or not to restart the instance when creating an AMI",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "Array of tags to assign to the created AMI",
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
			Description: "Tag key to assign to the created AMI",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "Tag value to assign to the created AMI",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"add_same_tag_to_snapshot": {
			Description: "Whether to add the tags assigned to the AMI to the EBS snapshot as well",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"recreate_image_if_ami_status_failed": {
			Description: "Whether or not to retry when a job fails",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
