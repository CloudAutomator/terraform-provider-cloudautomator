package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
