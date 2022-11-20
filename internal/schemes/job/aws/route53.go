package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UpdateRecordSetActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Description: "Host zone for updating the resource record set",
			Type:        schema.TypeString,
			Required:    true,
		},
		"record_set_name": {
			Description: "Resource record set to be updated",
			Type:        schema.TypeString,
			Required:    true,
		},
		"record_set_type": {
			Description: "Resource record type",
			Type:        schema.TypeString,
			Required:    true,
		},
		"record_set_value": {
			Description: "Resource Record Set Value",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
