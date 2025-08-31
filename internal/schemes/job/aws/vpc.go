package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CreateNatGatewayActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the NAT Gateway will be created",
			Type:        schema.TypeString,
			Required:    true,
		},
		"subnet_id": {
			Description: "The subnet ID in which to create the NAT Gateway",
			Type:        schema.TypeString,
			Required:    true,
		},
		"allocation_id": {
			Description: "The allocation ID of an Elastic IP address to associate with the NAT Gateway",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"connectivity_type": {
			Description: "Indicates whether the NAT Gateway supports public or private connectivity",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tags": {
			Description: "Array of tags to assign to the created NAT Gateway",
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
		"trace_status": {
			Description: "Whether to verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}