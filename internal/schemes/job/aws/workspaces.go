package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RebootWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "Tag key used to identify the target resource",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_value": {
			Description: "Tag value used to identify the target resource",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RebuildWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "Tag key used to identify the target resource",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_value": {
			Description: "Tag value used to identify the target resource",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func StartWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
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
	}
}

func TerminateWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_workspace": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
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
			Optional:    true,
		},
	}
}
