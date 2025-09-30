package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CreateNatGatewayActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region where the NAT Gateway will be created",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "Array of tags to be assigned to the NAT Gateway",
			Type:        schema.TypeSet,
			Required:    true,
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
		"allocation_id": {
			Description: "Allocation ID of the Elastic IP to assign to the NAT Gateway",
			Type:        schema.TypeString,
			Required:    true,
		},
		"nat_gateway_name": {
			Description: "Name tag value for the NAT Gateway",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"subnet_id": {
			Description: "Subnet ID where the NAT Gateway will be created",
			Type:        schema.TypeString,
			Required:    true,
		},
		"route_table_id": {
			Description: "Route table ID to add a route targeting the NAT Gateway",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}