package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DelayActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delay_minutes": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func NoActionActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}
