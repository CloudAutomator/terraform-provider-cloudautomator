package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ChangeElasticacheNodeTypeActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target ElastiCache cluster resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "Tag key used to identify the ElastiCache cluster",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_value": {
			Description: "Tag value used to identify the ElastiCache cluster",
			Type:        schema.TypeString,
			Required:    true,
		},
		"node_type": {
			Description: "Node type after modification",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
