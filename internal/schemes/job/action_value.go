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

func ChangeRdsClusterInstanceClassActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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
		"db_instance_class": {
			Description: "DB instance class after modification",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func ChangeRdsInstanceClassActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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
		"db_instance_class": {
			Description: "DB instance class after modification",
			Type:        schema.TypeString,
			Required:    true,
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

func CopyRdsClusterSnapshotActionValueFields() map[string]*schema.Schema {
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
		"specify_rds_cluster_snapshot": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_cluster_snapshot_id": {
			Description: "Target DB Cluster snapshot ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_rds_cluster_id": {
			Description: "Target DB Cluster ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"kms_key_id": {
			Description: "KMS key ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CopyRdsSnapshotActionValueFields() map[string]*schema.Schema {
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
		"specify_rds_snapshot": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_snapshot_id": {
			Description: "Target DB snapshot ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_rds_instance_id": {
			Description: "Target RDS instance ID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"option_group_name": {
			Description: "Option group name to be set for the destination region",
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

func CreateRdsClusterSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "Target DB Cluster ID",
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
		"db_cluster_snapshot_identifier": {
			Description: "Prefix for DB cluster snapshot to be created",
			Type:        schema.TypeString,
			Required:    true,
		},
		"generation": {
			Description: "Number of DB cluster snapshots to manage generations",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateRdsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_snapshot_id": {
			Description: "Target DB snapshot ID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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
			Description: "Number of DB snapshots generation management",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateRedshiftSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_cluster": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_snapshot_identifier": {
			Description: "Name to be set for the snapshot",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_identifier": {
			Description: "Target Cluster ID",
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
			Description: "Number of snapshots to be managed for generation",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DelayActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delay_minutes": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func DeleteClusterActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_identifier": {
			Description: "Target Redshift cluster ID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"final_cluster_snapshot_identifier": {
			Description: "Snapshot name to be taken when deleting Redshift cluster",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_final_cluster_snapshot": {
			Description: "Whether to skip taking snapshots when deleting Redshift clusters",
			Type:        schema.TypeString,
			Required:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DeleteRdsClusterActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "Target DB Cluster ID",
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
		"final_db_snapshot_identifier": {
			Description: "Snapshot name to be taken when deleting DB cluster",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_final_snapshot": {
			Description: "Whether to skip taking snapshots when deleting DB clusters",
			Type:        schema.TypeString,
			Required:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DeleteRdsInstanceActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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
		"final_rds_snapshot_id": {
			Description: "DB snapshot name to be retrieved when deleting RDS instance",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_final_rds_snapshot": {
			Description: "Whether to skip taking DB snapshots when deleting RDS instances",
			Type:        schema.TypeString,
			Required:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

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

func DescribeMetadataActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dr_configuration_id": {
			Description: "DR Configuration ID",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func DisasterRecoveryActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"describe_metadata_job_id": {
			Description: "Describe Metadata Job ID",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RebootRdsInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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

func RestoreFromClusterSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_identifier": {
			Description: "Redshift cluster ID after restore",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_identifier": {
			Description: "Redshift snapshot ID to be used for restore",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_parameter_group_name": {
			Description: "Parameter group name to be set for the restored Redshift cluster",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cluster_subnet_group_name": {
			Description: "Name of the subnet group where the restored Redshift cluster will be located",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"port": {
			Description: "Port number of the DB cluster after restoration",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"publicly_accessible": {
			Description: "Whether to make the restored Redshift cluster publicly accessible or not",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"availability_zone": {
			Description: "AvailabilityZone of the Redshift cluster after restore",
			Type:        schema.TypeString,
			Required:    true,
		},
		"vpc_security_group_ids": {
			Description: "Array containing the security group IDs to be set for the restored Redshift cluster",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"allow_version_upgrade": {
			Description: "Whether to enable automatic minor version upgrades on the Redshift cluster after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"delete_cluster_snapshot": {
			Description: "Whether to delete Redshift snapshots used for restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RestoreRdsClusterActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_instance_identifier": {
			Description: "ID of DB instance after restoration",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "ID of DB cluster after restore",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_identifier": {
			Description: "DB snapshot ID to be used for restore",
			Type:        schema.TypeString,
			Required:    true,
		},
		"engine": {
			Description: "DB engine of DB cluster after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"engine_version": {
			Description: "DB engine version of the DB cluster after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_instance_class": {
			Description: "DB instance class after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_subnet_group_name": {
			Description: "DB subnet group name where the DB cluster will be located after restoration",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"publicly_accessible": {
			Description: "Whether to make the DB cluster publicly accessible after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"availability_zone": {
			Description: "AZ to deploy DB cluster after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vpc_security_group_ids": {
			Description: "Array containing the security group IDs to be set for the restored DB cluster",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"port": {
			Description: "Port number of the DB cluster after restoration",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"db_cluster_parameter_group_name": {
			Description: "Parameter group name to be set for the DB cluster after restoration",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_parameter_group_name": {
			Description: "Parameter group name to be set for the DB instance after restoration",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"option_group_name": {
			Description: "Option group name to be set for the DB cluster after restoration",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"auto_minor_version_upgrade": {
			Description: "Whether to enable automatic minor version upgrades on the DB cluster after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"delete_db_cluster_snapshot": {
			Description: "Whether to delete DB snapshots used for restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RestoreRdsInstanceActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_snapshot_id": {
			Description: "Target DB snapshot ID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_engine": {
			Description: "DB engine of the RDS instance after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"license_model": {
			Description: "Licensing model of the RDS instance after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_instance_class": {
			Description: "DB instance class after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"multi_az": {
			Description: "Whether or not to configure the RDS instance in a Multi-AZ configuration after restoration",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"storage_type": {
			Description: "Storage type of the RDS instance after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"iops": {
			Description: "IOPS value of RDS instance after restore",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vpc": {
			Description: "ID of the VPC where the RDS instance will be placed after restoration",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"subnet_group": {
			Description: "DB subnet group name where the restored RDS instance will be located",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"publicly_accessible": {
			Description: "Whether to make the restored RDS instance publicly accessible or not",
			Type:        schema.TypeString,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"availability_zone": {
			Description: "AZ to deploy RDS instance after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vpc_security_group_ids": {
			Description: "Array containing the security group IDs to be set for the restored RDS instance",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"db_name": {
			Description: "Database name of the restored RDS instance",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"port": {
			Description: "Port number of the RDS instance after restoration",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"parameter_group": {
			Description: "Parameter group name to be set for the restored RDS instance",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"option_group": {
			Description: "Option group name to be set for the restored RDS instance",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"auto_minor_version_upgrade": {
			Description: "Whether to enable automatic minor version upgrade on the RDS instance after restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"delete_rds_snapshot": {
			Description: "Whether to delete DB snapshots used for restore",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_key": {
			Description: "Tag keys to assign to the restored RDS instance",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "Tag value to assign to the RDS instance after restoration",
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

func StartRdsClustersActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "Target DB Cluster ID",
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

func StartRdsInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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

func StopRdsClustersActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "Target DB Cluster ID",
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

func StopRdsInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "AWS Region in which the target resource resides",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "How to identify target resources",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "Target DB instance ID",
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
