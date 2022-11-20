package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeregisterInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "Target EC2 instance ID",
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
		"load_balancer_name": {
			Description: "Name of the ELB (CLB) to unregister the EC2 instance",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func DeregisterTargetInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"target_group_arn": {
			Description: "ARN of the target group",
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

func RegisterInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "Target EC2 instance ID",
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
		"load_balancer_name": {
			Description: "Name of the ELB (CLB) where the EC2 instance is registered",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RegisterTargetInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"target_group_arn": {
			Description: "ARN of the target group",
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
