package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func BulkDeleteRdsClusterSnapshotsActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"exclude_by_tag_bulk_delete_rds_cluster_snapshots": {
			Description: "Specifies whether to exclude DB cluster snapshots with certain tags from deletion",
			Type:        schema.TypeBool,
			Required:    true,
		},
		"exclude_by_tag_key_bulk_delete_rds_cluster_snapshots": {
			Description: "The tag key used to identify DB cluster snapshots to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"exclude_by_tag_value_bulk_delete_rds_cluster_snapshots": {
			Description: "The tag value used to identify DB cluster snapshots to exclude from deletion",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"specify_base_date": {
			Description: "Specifies the method for determining which DB cluster snapshots to delete",
			Type:        schema.TypeString,
			Required:    true,
		},
		"before_days": {
			Description: "The number of days used to identify DB cluster snapshots to be deleted",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"before_date": {
			Description: "The date used to identify DB cluster snapshots for deletion",
			Type:        schema.TypeString,
			Optional:    true,
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
