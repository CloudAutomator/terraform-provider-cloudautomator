package cloudautomator

import (
	"context"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"
	schemes "terraform-provider-cloudautomator/internal/schemes/job"
	aws "terraform-provider-cloudautomator/internal/schemes/job/aws"
	gcp "terraform-provider-cloudautomator/internal/schemes/job/gcp"
	other "terraform-provider-cloudautomator/internal/schemes/job/other"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJobRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Job ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Job Name",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"group_id": {
				Description: "Group ID",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"aws_account_id": {
				Description: "AWS account ID",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"aws_account_ids": {
				Description: "AWS account IDs",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"rule_type": {
				Description: "Trigger type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cron_rule_value": {
				Description: "Timer trigger value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CronRuleValueFields(),
				},
			},
			"schedule_rule_value": {
				Description: "Schedule trigger value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.ScheduleRuleValueFields(),
				},
			},
			"sqs_v2_rule_value": {
				Description: "SQS trigger value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.SqsV2RuleValueFields(),
				},
			},
			"action_type": {
				Description: "Action type",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"attach_user_policy_action_value": {
				Description: "\"IAM: Attach Policy to IAM User\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.AttachUserPolicyActionValueFields(),
				},
			},
			"authorize_security_group_ingress_action_value": {
				Description: "\"EC2: Authorize security group ingress\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.AuthorizeSecurityGroupIngressyActionValueFields(),
				},
			},
			"bulk_delete_ebs_snapshots_action_value": {
				Description: "\"EC2: Delete old EBS Snapshots\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.BulkDeleteEBSSnapshotsActionValueFields(),
				},
			},
			"bulk_delete_images_action_value": {
				Description: "\"EC2: Delete old AMIs and Snapshots\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.BulkDeleteImagesActionValueFields(),
				},
			},
			"bulk_delete_rds_cluster_snapshots_action_value": {
				Description: "\"RDS(Aurora): Delete old DB cluster snapshots\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.BulkDeleteRdsClusterSnapshotsActionValueFields(),
				},
			},
			"bulk_stop_instances_action_value": {
				Description: "\"EC2: Stop ALL instances\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.BulkStopInstancesActionValueFields(),
				},
			},
			"change_rds_cluster_instance_class_action_value": {
				Description: "\"RDS(Aurora): Change DB instance class\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.ChangeRdsClusterInstanceClassActionValueFields(),
				},
			},
			"change_rds_instance_class_action_value": {
				Description: "\"RDS: Change DB instance class\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.ChangeRdsInstanceClassActionValueFields(),
				},
			},
			"change_instance_type_action_value": {
				Description: "\"EC2: Change instance type\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.ChangeInstanceTypeActionValueFields(),
				},
			},
			"copy_ebs_snapshot_action_value": {
				Description: "\"EC2: Copy EBS snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CopyEbsSnapshotActionValueFields(),
				},
			},
			"copy_image_action_value": {
				Description: "\"EC2: Copy AMI\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CopyImageActionValueFields(),
				},
			},
			"copy_rds_cluster_snapshot_action_value": {
				Description: "\"RDS(Aurora): Copy DB cluster snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CopyRdsClusterSnapshotActionValueFields(),
				},
			},
			"copy_rds_snapshot_action_value": {
				Description: "\"RDS: Copy DB snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CopyRdsSnapshotActionValueFields(),
				},
			},
			"create_fsx_backup_action_value": {
				Description: "\"FSx: Create a backup\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateFSxBackupActionValueFields(),
				},
			},
			"create_ebs_snapshot_action_value": {
				Description: "\"EC2: Create EBS snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateEbsSnapshotActionValueFields(),
				},
			},
			"create_image_action_value": {
				Description: "\"EC2: Create AMI\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateImageActionValueFields(),
			},
				},
			"create_nat_gateway_action_value": {
				Description: "\"VPC: Create NAT Gateway\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateNatGatewayActionValueFields(),
				},
			},
			},
			"create_rds_cluster_snapshot_action_value": {
				Description: "\"RDS(Aurora): Create DB cluster snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateRdsClusterSnapshotActionValueFields(),
				},
			},
			"create_rds_snapshot_action_value": {
				Description: "\"RDS: Create DB snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateRdsSnapshotActionValueFields(),
				},
			},
			"create_redshift_snapshot_action_value": {
				Description: "\"Redshift: Create cluster snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.CreateRedshiftSnapshotActionValueFields(),
				},
			},
			"delay_action_value": {
				Description: "\"Other: Delay\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: other.DelayActionValueFields(),
				},
			},
			"delete_cluster_action_value": {
				Description: "\"Redshift: Delete cluster\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DeleteClusterActionValueFields(),
				},
			},
			"delete_rds_cluster_action_value": {
				Description: "\"RDS(Aurora): Delete DB cluster\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DeleteRdsClusterActionValueFields(),
				},
			},
			"delete_rds_instance_action_value": {
				Description: "\"RDS: Delete DB instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DeleteRdsInstanceActionValueFields(),
				},
			},
			"deregister_instances_action_value": {
				Description: "\"ELB(CLB): De-register EC2 instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DeregisterInstancesActionValueFields(),
				},
			},
			"deregister_target_instances_action_value": {
				Description: "\"ELB(ALB/NLB): Deregister EC2 instances from target group\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DeregisterTargetInstancesActionValueFields(),
				},
			},
			"describe_metadata_action_value": {
				Description: "\"DR: Update EC2 instance metadata\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DescribeMetadataActionValueFields(),
				},
			},
			"detach_user_policy_action_value": {
				Description: "\"IAM: Detach Policy to IAM User\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DetachUserPolicyActionValueFields(),
				},
			},
			"disaster_recovery_action_value": {
				Description: "\"DR: Launch EC2 instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DisasterRecoveryActionValueFields(),
				},
			},
			"dynamodb_start_backup_job_action_value": {
				Description: "\"DynamoDB: Backup table\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DynamodbStartBackupJobActionValueFields(),
				},
			},
			"ec2_start_backup_job_action_value": {
				Description: "\"EC2: Backup instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.Ec2StartBackupJobActionValueFields(),
				},
			},
			"efs_start_backup_job_action_value": {
				Description: "\"EFS: Backup file system\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.EfsStartBackupJobActionValueFields(),
				},
			},
			"google_compute_insert_machine_image_action_value": {
				Description: "\"Compute Engine: create machine image\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: gcp.GoogleComputeInsertMachineImageActionValueFields(),
				},
			},
			"reboot_rds_instances_action_value": {
				Description: "\"RDS: Reboot DB instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RebootRdsInstancesActionValueFields(),
				},
			},
			"reboot_workspaces_action_value": {
				Description: "\"WorkSpaces: Reboot WorkSpace\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RebootWorkspacesActionValueFields(),
				},
			},
			"rebuild_workspaces_action_value": {
				Description: "\"WorkSpaces: Rebuild WorkSpace\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RebuildWorkspacesActionValueFields(),
				},
			},
			"register_instances_action_value": {
				Description: "\"ELB(CLB): Register EC2 instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RegisterInstancesActionValueFields(),
				},
			},
			"register_target_instances_action_value": {
				Description: "\"ELB(ALB/NLB): Register EC2 instances to target group\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RegisterTargetInstancesActionValueFields(),
				},
			},
			"restore_from_cluster_snapshot_action_value": {
				Description: "\"Redshift: Restore from snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RestoreFromClusterSnapshotActionValueFields(),
				},
			},
			"restore_rds_cluster_action_value": {
				Description: "\"RDS(Aurora): Restore DB cluster from DB cluster snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RestoreRdsClusterActionValueFields(),
				},
			},
			"restore_rds_instance_action_value": {
				Description: "\"RDS: Restore from DB snapshot\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RestoreRdsInstanceActionValueFields(),
				},
			},
			"revoke_security_group_ingress_action_value": {
				Description: "\"EC2: Revoke security group ingress\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RevokeSecurityGroupIngressActionValueFields(),
				},
			},
			"run_ecs_tasks_fargate_action_value": {
				Description: "\"ECS: Run task (Fargate)\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.RunEcsTasksFargateActionValueFields(),
				},
			},
			"s3_start_backup_job_action_value": {
				Description: "\"S3: Backup\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.S3StartBackupJobActionValueFields(),
				},
			},
			"send_command_action_value": {
				Description: "\"EC2: Send command on instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.SendCommandActionValueFields(),
				},
			},
			"start_instances_action_value": {
				Description: "\"EC2: Start instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StartInstancesActionValueFields(),
				},
			},
			"start_rds_clusters_action_value": {
				Description: "\"RDS(Aurora): Start DB cluster\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StartRdsClustersActionValueFields(),
				},
			},
			"start_rds_instances_action_value": {
				Description: "\"RDS: Start DB instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StartRdsInstancesActionValueFields(),
				},
			},
			"stop_ecs_tasks_action_value": {
				Description: "\"ECS: Stop task\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StopEcsTasksActionValueFields(),
				},
			},
			"stop_instances_action_value": {
				Description: "\"EC2: Stop instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StopInstancesActionValueFields(),
				},
			},
			"stop_rds_clusters_action_value": {
				Description: "\"RDS(Aurora): Stop DB cluster\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StopRdsClustersActionValueFields(),
				},
			},
			"stop_rds_instances_action_value": {
				Description: "\"RDS: Stop DB instance\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StopRdsInstancesActionValueFields(),
				},
			},
			"start_workspaces_action_value": {
				Description: "\"WorkSpaces: Start WorkSpace\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.StartWorkspacesActionValueFields(),
				},
			},
			"terminate_workspaces_action_value": {
				Description: "\"WorkSpaces: Remove WorkSpace\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.TerminateWorkspacesActionValueFields(),
				},
			},
			"update_record_set_action_value": {
				Description: "\"Route 53: Update Resource Record Set\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.UpdateRecordSetActionValueFields(),
				},
			},
			"windows_update_action_value": {
				Description: "\"EC2: Windows Update to instance (Old version)\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.WindowsUpdateActionValueFields(),
				},
			},
			"windows_update_v2_action_value": {
				Description: "\"EC2: Windows Update to instance (New version)\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.WindowsUpdateV2ActionValueFields(),
				},
			},
			"allow_runtime_action_values": {
				Description: "Whether the value of the action setting is specified at runtime or not",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"effective_date": {
				Description: "Effective date",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"expiration_date": {
				Description: "Expiration date",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"completed_post_process_id": {
				Description: "Array containing post-process IDs to be executed if the job is successful",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"failed_post_process_id": {
				Description: "Array containing post-process IDs to be executed if the job fails",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceJobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Get("id").(string)

	job, _, err := c.GetJob(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("id", job.Id)
	d.Set("name", job.Name)
	d.Set("group_id", job.GroupId)
	d.Set("aws_account_id", job.AwsAccountId)
	d.Set("aws_account_ids", utils.FlattenIntList(job.AwsAccountIds))

	d.Set("rule_type", job.RuleType)

	switch job.RuleType {
	case "cron", "schedule", "sqs_v2":
		ruleValueBlockName := fmt.Sprintf("%s_rule_value", job.RuleType)
		if err := d.Set(ruleValueBlockName, []interface{}{job.RuleValue}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}
	d.Set("action_type", job.ActionType)
	actionValueBlockName := fmt.Sprintf("%s_action_value", job.ActionType)
	if err := d.Set(actionValueBlockName, []interface{}{job.ActionValue}); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("allow_runtime_action_values", job.AllowRuntimeActionValues)
	d.Set("effective_date", job.EffectiveDate)
	d.Set("expiration_date", job.ExpirationDate)
	d.Set("completed_post_process_id", utils.FlattenIntList(job.CompletedPostProcessId))
	d.Set("failed_post_process_id", utils.FlattenIntList(job.FailedPostProcessId))

	d.SetId(job.Id)

	return diags
}
