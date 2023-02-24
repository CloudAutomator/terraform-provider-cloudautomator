package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func InvokeLambdaFunctionActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"function_name": {
			Description: "Function name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"payload": {
			Description: "Event JSON",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
