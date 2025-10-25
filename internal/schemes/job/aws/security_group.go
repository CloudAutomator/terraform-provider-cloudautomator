package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthorizeSecurityGroupIngressActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_security_group": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"security_group_id": {
			Description: "Target security group ID",
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
		"ip_protocol": {
			Description: "IP protocol",
			Type:        schema.TypeString,
			Required:    true,
		},
		"to_port": {
			Description: "port",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cidr_ip": {
			Description: "CIDR IP",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RevokeSecurityGroupIngressActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_security_group": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"security_group_id": {
			Description: "Target security group ID",
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
		"ip_protocol": {
			Description: "IP protocol",
			Type:        schema.TypeString,
			Required:    true,
		},
		"to_port": {
			Description: "port",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cidr_ip": {
			Description: "CIDR IP",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}
