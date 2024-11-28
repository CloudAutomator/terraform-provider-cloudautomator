package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthorizeSecurityGroupIngressyActionValueFields() map[string]*schema.Schema {
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

func BulkDeleteEBSSnapshotsActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_by_tag_bulk_delete_ebs_snapshots": {
			Description: "Specifies whether to exclude EBS Snapshots with certain tags from deletion",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"exclude_by_tag_key_bulk_delete_ebs_snapshots": {
			Description: "The tag key used to identify EBS Snapshots to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"exclude_by_tag_value_bulk_delete_ebs_snapshots": {
			Description: "The tag value used to identify EBS Snapshots to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"specify_base_date": {
			Description: "Specifies the method for determining which EBS Snapshots to delete",
			Type:        schema.TypeString,
			Required:    true,
		},
		"before_days": {
			Description: "The number of days used to identify EBS Snapshots to be deleted",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"before_date": {
			Description: "The date used to identify EBS Snapshots for deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func BulkDeleteImagesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_by_tag_bulk_delete_images": {
			Description: "Specifies whether to exclude AMIs with certain tags from deletion",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"exclude_by_tag_key_bulk_delete_images": {
			Description: "The tag key used to identify AMIs to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"exclude_by_tag_value_bulk_delete_images": {
			Description: "The tag value used to identify AMIs to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"specify_base_date": {
			Description: "Specifies the method for determining which AMIs to delete",
			Type:        schema.TypeString,
			Required:    true,
		},
		"before_days": {
			Description: "The number of days used to identify AMIs to be deleted",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"before_date": {
			Description: "The date used to identify AMIs for deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

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

func CopyEbsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "AWS Region from which the copy was made",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "AWS Region to copy to",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_ebs_snapshot": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_id": {
			Description: "Target EBS snapshot ID",
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
			Required:    true,
		},
	}
}

func CopyImageActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "AWS Region from which the copy was made",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "AWS Region to copy to",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_image": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"source_image_id": {
			Description: "対象のAMIのID",
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
			Required:    true,
		},
	}
}

func CreateEbsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_volume": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"volume_id": {
			Description: "対象のEBSボリュームID",
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
		"generation": {
			Description: "Number of EBS volumes to manage generation",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"description": {
			Description: "Description to be set for EBS volume",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tags": {
			Description: "Array of tags to assign to the created EBS volume",
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
		"additional_tag_key": {
			Description: "Tag key to assign to the created EBS volume",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "Tag key to assign to the created EBS volume Tag value to assign to the created EBS volume",
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

func CreateImageActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_image_instance": {
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
		"generation": {
			Description: "Number of AMI generations to be managed",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"image_name": {
			Description: "Image name to set for AMI",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description to be set in AMI",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"reboot_instance": {
			Description: "Flag whether or not to restart the instance when creating an AMI",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "Array of tags to assign to the created AMI",
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
		"additional_tag_key": {
			Description: "Tag key to assign to the created AMI",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "Tag value to assign to the created AMI",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"add_same_tag_to_snapshot": {
			Description: "Whether to add the tags assigned to the AMI to the EBS snapshot as well",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"recreate_image_if_ami_status_failed": {
			Description: "Whether or not to retry when a job fails",
			Type:        schema.TypeString,
			Optional:    true,
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
