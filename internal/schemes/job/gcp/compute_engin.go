package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GoogleComputeInsertMachineImageActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "GCP Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_vm_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"vm_instance_label_key": {
			Description: "label key used to identify the target resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vm_instance_label_value": {
			Description: "label value used to identify the target resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vm_instance_id": {
			Description: "VM Instance ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"machine_image_basename": {
			Description: "Name of the machine image to created",
			Type:        schema.TypeString,
			Required:    true,
		},
		"machine_image_storage_location": {
			Description: "Machine image storage location",
			Type:        schema.TypeString,
			Required:    true,
		},
		"generation": {
			Description: "Tag value used to identify the target resource",
			Type:        schema.TypeString,
			Required:    true,
		},
		"project_id": {
			Description: "Project ID to which the target VM instance belongs",
			Type:        schema.TypeString,
			Required:    true,
		},
		"machine_image_description": {
			Description: "Description to set for machine image to created",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
