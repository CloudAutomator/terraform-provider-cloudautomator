package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BulkStopInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_by_tag": {
			Description: "Whether to exclude instances with the specified tag from the target",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"exclude_by_tag_key": {
			Description: "Tag key used to exclude instances from the target",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"exclude_by_tag_value": {
			Description: "Tag value used to exclude instances from the target",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func ChangeInstanceTypeActionValueFields() map[string]*schema.Schema {
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
		"instance_type": {
			Description: "Changed instance type",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func Ec2StartBackupJobActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region",
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
		"backup_vault_name": {
			Description: "Backup Vault Name",
			Type:        schema.TypeString,
			Required:    true,
		},
		"lifecycle_delete_after_days": {
			Description: "Number of days to hold recovery point",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"iam_role_arn": {
			Description: "IAM Role ARN",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "Array of tags to be added to the recovery point",
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
	}
}

func SendCommandActionValueFields() map[string]*schema.Schema {
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
		"command": {
			Description: "Commands to execute",
			Type:        schema.TypeString,
			Required:    true,
		},
		"comment": {
			Description: "Comments to be set for the command",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"document_name": {
			Description: "Command Type",
			Type:        schema.TypeString,
			Required:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"output_s3_bucket_name": {
			Description: "Name of the S3 bucket in which to store the execution results",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_key_prefix": {
			Description: "S3 prefix to store execution results",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"timeout_seconds": {
			Description: "Instance connection timeout time (seconds)",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"execution_timeout_seconds": {
			Description: "Timeout period for command execution (seconds)",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StartInstancesActionValueFields() map[string]*schema.Schema {
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
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status_checks_enable": {
			Description: "Whether status check is performed or not. Can be `true` only if `true` is specified in trace_status",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StopInstancesActionValueFields() map[string]*schema.Schema {
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
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func WindowsUpdateActionValueFields() map[string]*schema.Schema {
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
		"comment": {
			Description: "Comments to be set for the command",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"document_name": {
			Description: "`AWS-InstallMissingWindowsUpdates` 固定",
			Type:        schema.TypeString,
			Required:    true,
		},
		"kb_article_ids": {
			Description: "Array containing KB to be excluded",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_bucket_name": {
			Description: "Name of the S3 bucket in which to store the execution results",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_key_prefix": {
			Description: "S3 prefix to store execution results",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"update_level": {
			Description: "Update Level",
			Type:        schema.TypeString,
			Required:    true,
		},
		"timeout_seconds": {
			Description: "Timeout time (sec)",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func WindowsUpdateV2ActionValueFields() map[string]*schema.Schema {
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
		"allow_reboot": {
			Description: "Whether to allow reboots caused by applying Windows Update",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_severity": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"severity_levels": {
			Description: "Severity of Windows Update to be applied",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"output_s3_bucket_name": {
			Description: "Name of the S3 bucket in which to store the execution results",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_key_prefix": {
			Description: "S3 prefix to store execution results",
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
