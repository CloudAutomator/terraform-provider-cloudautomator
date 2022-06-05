package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AuthorizeSecurityGroupIngressyActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_security_group": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"security_group_id": {
			Description: "対象のセキュリティグループID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ip_protocol": {
			Description: "通信プロトコル",
			Type:        schema.TypeString,
			Required:    true,
		},
		"to_port": {
			Description: "ポート番号",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cidr_ip": {
			Description: "送信元IPのCIDRアドレス",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func ChangeRdsClusterInstanceClassActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_instance_class": {
			Description: "変更後のDBインスタンスクラス",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func ChangeRdsInstanceClassActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_instance_class": {
			Description: "変更後のDBインスタンスクラス",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func ChangeInstanceTypeActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"instance_type": {
			Description: "変更後のインスタンスタイプ",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func CopyEbsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "コピー元のAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "コピー先のAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_ebs_snapshot": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_id": {
			Description: "対象のEBSスナップショットID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "EBSスナップショットのコピー完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func CopyImageActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "コピー元のAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "コピー先のAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_image": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"source_image_id": {
			Description: "対象のAMIのID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "AMIのコピー完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func CopyRdsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source_region": {
			Description: "コピー元のAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"destination_region": {
			Description: "コピー先のAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_snapshot": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_snapshot_id": {
			Description: "対象のDBスナップショットID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_rds_instance_id": {
			Description: "対象のRDSインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"option_group_name": {
			Description: "コピー先リージョンに設定するオプショングループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "DBスナップショットのコピー完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func CreateEbsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_volume": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"volume_id": {
			Description: "対象のEBSボリュームID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"generation": {
			Description: "EBSボリュームの世代管理を行う数",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"description": {
			Description: "EBSボリュームに設定する説明",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_key": {
			Description: "作成したEBSボリュームに割り当てるタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "作成したEBSボリュームに割り当てるタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "EBSボリュームの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateImageActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_image_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"generation": {
			Description: "AMIの世代管理を行う数",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"image_name": {
			Description: "AMIに設定するイメージ名",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "AMIに設定する説明",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"reboot_instance": {
			Description: "AMI作成時にインスタンスを再起動するか否かのフラグ",
			Type:        schema.TypeString,
			Required:    true,
		},
		"additional_tags": {
			Description: "作成したAMIに割り当てるタグの配列",
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
			Description: "作成したAMIに割り当てるタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "作成したAMIに割り当てるタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"add_same_tag_to_snapshot": {
			Description: "AMIに割り当てたタグをEBSスナップショットにも追加するか否かのフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "AMIの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"recreate_image_if_ami_status_failed": {
			Description: "ジョブ失敗時にリトライを行うか否かのフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateRdsClusterSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "DBクラスターの特定に利用するID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_cluster_snapshot_identifier": {
			Description: "作成するDBクラスタースナップショットのプレフィックス",
			Type:        schema.TypeString,
			Required:    true,
		},
		"generation": {
			Description: "DBクラスタースナップショットの世代管理を行う数",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"trace_status": {
			Description: "DBクラスタースナップショットの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateRdsSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_snapshot_id": {
			Description: "対象のDBスナップショットID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"generation": {
			Description: "DBスナップショットの世代管理を行う数",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"trace_status": {
			Description: "DBスナップショットの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func CreateRedshiftSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_cluster": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_snapshot_identifier": {
			Description: "スナップショットに設定する名前",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_identifier": {
			Description: "対象のクラスターID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"generation": {
			Description: "スナップショットの世代管理を行う数",
			Type:        schema.TypeInt,
			Required:    true,
		},
		"trace_status": {
			Description: "スナップショットの作成完了をジョブ完了の判定にするフラグ",
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
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_identifier": {
			Description: "対象のRedshiftクラスターのidentifier",
			Type:        schema.TypeString,
			Required:    true,
		},
		"final_cluster_snapshot_identifier": {
			Description: "Redshiftクラスター削除時に取得するRedshiftクラスタースナップショット名",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_final_cluster_snapshot": {
			Description: "Redshiftクラスター削除時のRedshiftクラスタースナップショット取得をスキップするか",
			Type:        schema.TypeString,
			Required:    true,
		},
		"trace_status": {
			Description: "Redshiftクラスターの削除完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DeleteRdsClusterActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "DBクラスターの特定に利用するID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"final_db_snapshot_identifier": {
			Description: "DBクラスター削除時に取得するDBクラスタースナップショット名",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_final_snapshot": {
			Description: "DBクラスター削除時のDBクラスタースナップショット取得をスキップするか",
			Type:        schema.TypeString,
			Required:    true,
		},
		"trace_status": {
			Description: "DBクラスターの削除完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DeleteRdsInstanceActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"final_rds_snapshot_id": {
			Description: "RDSインスタンス削除時に取得するDBスナップショット名",
			Type:        schema.TypeString,
			Required:    true,
		},
		"skip_final_rds_snapshot": {
			Description: "RDSインスタンス削除時のDBスナップショット取得をスキップするか",
			Type:        schema.TypeString,
			Required:    true,
		},
		"trace_status": {
			Description: "DBスナップショットの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func DeregisterInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"load_balancer_name": {
			Description: "EC2インスタンスを登録解除するELB(CLB)名",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func DeregisterTargetInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"target_group_arn": {
			Description: "対象のターゲットグループのARN",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RebootRdsInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RebootWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RegisterInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"load_balancer_name": {
			Description: "EC2インスタンスを登録するELB(CLB)名",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RegisterTargetInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"target_group_arn": {
			Description: "対象のターゲットグループのARN",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func RestoreFromClusterSnapshotActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_identifier": {
			Description: "リストア後のRedshiftクラスターのidentifier",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_identifier": {
			Description: "リストアに使用するRedshiftスナップショットID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cluster_parameter_group_name": {
			Description: "リストア後のRedshiftクラスターに設定するパラメータグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cluster_subnet_group_name": {
			Description: "リストア後のRedshiftクラスターを配置するサブネットグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"port": {
			Description: "リストア後のDBクラスターの接続ポート番号",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"publicly_accessible": {
			Description: "リストア後のRedshiftクラスターをパブリックアクセス可能にするか否か",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"availability_zone": {
			Description: "リストア後のRedshiftクラスターのAvailabilityZone",
			Type:        schema.TypeString,
			Required:    true,
		},
		"vpc_security_group_ids": {
			Description: "リストア後のRedshiftクラスターに設定するセキュリティグループIDが含まれる配列",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"allow_version_upgrade": {
			Description: "リストア後のRedshiftクラスターで自動マイナーバージョンアップグレードを有効にするかどうか",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"delete_cluster_snapshot": {
			Description: "リストアに利用したRedshiftスナップショットを削除するかどうか",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RestoreRdsClusterActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_instance_identifier": {
			Description: "リストア後のDBインスタンスのidentifier",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "リストア後のDBクラスターのidentifier",
			Type:        schema.TypeString,
			Required:    true,
		},
		"snapshot_identifier": {
			Description: "リストアに使用するDBスナップショットID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"engine": {
			Description: "リストア後のDBクラスターのDBエンジン",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"engine_version": {
			Description: "リストア後のDBクラスターのDBエンジンのバージョン",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_instance_class": {
			Description: "リストア後のDBインスタンスクラス",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_subnet_group_name": {
			Description: "リストア後のDBクラスターを配置するDBサブネットグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"publicly_accessible": {
			Description: "リストア後のDBクラスターをパブリックアクセス可能にするか否か",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"availability_zone": {
			Description: "リストア後のDBクラスターを配置するAZ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vpc_security_group_ids": {
			Description: "リストア後のDBクラスターに設定するセキュリティグループIDが含まれる配列",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"port": {
			Description: "リストア後のDBクラスターの接続ポート番号",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"db_cluster_parameter_group_name": {
			Description: "リストア後のDBクラスターに設定するパラメータグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_parameter_group_name": {
			Description: "リストア後のDBインスタンスに設定するパラメータグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"option_group_name": {
			Description: "リストア後のDBクラスターに設定するオプショングループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"auto_minor_version_upgrade": {
			Description: "リストア後のDBクラスターで自動マイナーバージョンアップグレードを有効にするかどうか",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"delete_db_cluster_snapshot": {
			Description: "リストアに利用したDBスナップショットを削除するかどうか",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RestoreRdsInstanceActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_snapshot_id": {
			Description: "対象のDBスナップショットID",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_engine": {
			Description: "リストア後のRDSインスタンスのDBエンジン",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"license_model": {
			Description: "リストア後のRDSインスタンスのライセンスモデル",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"db_instance_class": {
			Description: "リストア後のDBインスタンスクラス",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"multi_az": {
			Description: "リストア後のRDSインスタンスをMulti-AZ構成にするか否か",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"storage_type": {
			Description: "リストア後のRDSインスタンスのストレージタイプ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"iops": {
			Description: "リストア後のRDSインスタンスのIOPS値",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"vpc": {
			Description: "リストア後のRDSインスタンスを配置するVPCのID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"subnet_group": {
			Description: "リストア後のRDSインスタンスを配置するDBサブネットグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"publicly_accessible": {
			Description: "リストア後のRDSインスタンスをパブリックアクセス可能にするか否か",
			Type:        schema.TypeString,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"availability_zone": {
			Description: "リストア後のRDSインスタンスを配置するAZ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"vpc_security_group_ids": {
			Description: "リストア後のRDSインスタンスに設定するセキュリティグループIDが含まれる配列",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"db_name": {
			Description: "リストア後のRDSインスタンスのデータベース名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"port": {
			Description: "リストア後のRDSインスタンスの接続ポート番号",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"parameter_group": {
			Description: "リストア後のRDSインスタンスに設定するパラメータグループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"option_group": {
			Description: "リストア後のRDSインスタンスに設定するオプショングループ名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"auto_minor_version_upgrade": {
			Description: "リストア後のRDSインスタンスで自動マイナーバージョンアップグレードを有効にするかどうか",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"delete_rds_snapshot": {
			Description: "リストアに利用したDBスナップショットを削除するかどうか",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_key": {
			Description: "リストア後のRDSインスタンスに割り当てるタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"additional_tag_value": {
			Description: "リストア後のRDSインスタンスの割り当てるタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "RDSインスタンスの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func RevokeSecurityGroupIngressActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_security_group": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"security_group_id": {
			Description: "対象のセキュリティグループID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"ip_protocol": {
			Description: "通信プロトコル",
			Type:        schema.TypeString,
			Required:    true,
		},
		"to_port": {
			Description: "ポート番号",
			Type:        schema.TypeString,
			Required:    true,
		},
		"cidr_ip": {
			Description: "送信元IPのCIDRアドレス",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func SendCommandActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"command": {
			Description: "実行するコマンド",
			Type:        schema.TypeString,
			Required:    true,
		},
		"comment": {
			Description: "コマンドに設定するコメント",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"document_name": {
			Description: "コマンドの種類",
			Type:        schema.TypeString,
			Required:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"output_s3_bucket_name": {
			Description: "実行結果を保存するS3のバケット名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_key_prefix": {
			Description: "実行結果を保存するS3のプレフィックス",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "実行コマンドの終了ステータスをジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"timeout_seconds": {
			Description: "インスタンス接続のタイムアウト時間(秒)",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"execution_timeout_seconds": {
			Description: "コマンド実行のタイムアウト時間(秒)",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func StartInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "インスタンスの起動完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"status_checks_enable": {
			Description: "ステータスチェックを行うかどうか。 trace_statusで `true` を指定した場合のみ `true` を指定可能",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StartRdsClustersActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "対象のDBクラスターID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "DBクラスターの起動完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StartRdsInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "RDSインスタンスの起動完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StopInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "インスタンスの停止完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StopRdsClustersActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_cluster": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"db_cluster_identifier": {
			Description: "対象のDBクラスターID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "DBクラスターの停止完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StopRdsInstancesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_rds_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"rds_instance_id": {
			Description: "対象のDBインスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "RDSインスタンスの停止完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func StartWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func TerminateWorkspacesActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_workspace": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "WorkSpaceの作成完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func UpdateRecordSetActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_name": {
			Description: "リソースレコードセットを更新するホストゾーン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"record_set_name": {
			Description: "更新対象のリソースレコードセット",
			Type:        schema.TypeString,
			Required:    true,
		},
		"record_set_type": {
			Description: "リソースレコードタイプ",
			Type:        schema.TypeString,
			Required:    true,
		},
		"record_set_value": {
			Description: "リソースレコードセットの値",
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func WindowsUpdateActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"comment": {
			Description: "コマンドに設定するコメント",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"document_name": {
			Description: "`AWS-InstallMissingWindowsUpdates` 固定",
			Type:        schema.TypeString,
			Required:    true,
		},
		"kb_article_ids": {
			Description: "除外するKBが含まれた配列",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_bucket_name": {
			Description: "実行結果を保存するS3のバケット名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_key_prefix": {
			Description: "実行結果を保存するS3のプレフィックス",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"update_level": {
			Description: "アップデートレベル",
			Type:        schema.TypeString,
			Required:    true,
		},
		"timeout_seconds": {
			Description: "タイムアウト時間(秒)",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func WindowsUpdateV2ActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"region": {
			Description: "対象のリソースが存在するAWSリージョン",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_instance": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"instance_id": {
			Description: "対象のEC2インスタンスID",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_key": {
			Description: "対象リソースの特定に利用するタグのキー",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tag_value": {
			Description: "対象リソースの特定に利用するタグの値",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"allow_reboot": {
			Description: "Windows Updateの適用で発生する再起動を許容するか",
			Type:        schema.TypeString,
			Required:    true,
		},
		"specify_severity": {
			Description: "対象リソースを特定する方法",
			Type:        schema.TypeString,
			Required:    true,
		},
		"severity_levels": {
			Description: "適用するWindows Updateの重要度",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"output_s3_bucket_name": {
			Description: "実行結果を保存するS3のバケット名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"output_s3_key_prefix": {
			Description: "実行結果を保存するS3のバケット名",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"trace_status": {
			Description: "Windows Update完了をジョブ完了の判定にするフラグ",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
