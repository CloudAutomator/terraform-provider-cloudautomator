package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
