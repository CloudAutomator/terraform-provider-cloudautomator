package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GoogleComputeInsertMachineImageActionValueFields returns the schema for creating a machine image
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
			Description: "Number of machine image generations",
			Type:        schema.TypeInt,
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

// GoogleComputeStartVmInstancesActionValueFields returns the schema for starting VM instances
func GoogleComputeStartVmInstancesActionValueFields() map[string]*schema.Schema {
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
		"project_id": {
			Description: "Project ID to which the target VM instance belongs",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// GoogleComputeStopVmInstancesActionValueFields returns the schema for stopping VM instances
func GoogleComputeStopVmInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "GCP Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_vm_instance": {
			Description: "How to specify VM instance",
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
		"project_id": {
			Description: "Project ID to which the target VM instance belongs",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
