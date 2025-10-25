package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func RunEcsTasksFargateActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ecs_cluster": {
			Description: "ECS cluster name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"platform_version": {
			Description: "Platform version",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ecs_task_definition_family": {
			Description: "ECS task definition family",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ecs_task_count": {
			Description: "Number of ECS tasks to run",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"propagate_tags": {
			Description: "Propagate tags",
			Type:        schema.TypeString,
			Required:    true,
		},
		"enable_ecs_managed_tags": {
			Description: "Enable ECS managed tags",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"ecs_awsvpc_vpc": {
			Description: "ECS awsvpc vpc",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ecs_awsvpc_subnets": {
			Description: "ECS awsvpc subnets",
			Type:        schema.TypeList,
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"ecs_awsvpc_security_groups": {
			Description: "ECS awsvpc security groups",
			Type:        schema.TypeList,
			Required:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"ecs_awsvpc_assign_public_ip": {
			Description: "ECS awsvpc assign public ip",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func StopEcsTasksActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ecs_cluster": {
			Description: "Target ECS cluster name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_ecs_task": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"ecs_task_definition_family": {
			Description: "ECS task definition family name",
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
	}
}
