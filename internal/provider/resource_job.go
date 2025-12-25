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

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Description:   "Manage Cloud Automator job resources",
		CreateContext: resourceJobCreate,
		ReadContext:   resourceJobRead,
		UpdateContext: resourceJobUpdate,
		DeleteContext: resourceJobDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Job ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Job Name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"for_workflow": {
				Description: "for workflow",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"group_id": {
				Description: "Group ID",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"aws_account_id": {
				Description: "AWS account ID",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"aws_account_ids": {
				Description: "AWS account IDs",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"google_cloud_account_id": {
				Description: "Google Cloud account ID",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"rule_type": {
				Description: "Trigger type",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"cron_rule_value": {
				Description: "Timer trigger value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.CronRuleValueFields(),
				},
			},
			"schedule_rule_value": {
				Description: "Schedule trigger value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.ScheduleRuleValueFields(),
				},
			},
			"sqs_v2_rule_value": {
				Description: "SQS trigger value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.SqsV2RuleValueFields(),
				},
			},
			"action_type": {
				Description: "Action type",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
					Schema: aws.AuthorizeSecurityGroupIngressActionValueFields(),
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
			"bulk_delete_rds_snapshots_action_value": {
				Description: "\"RDS: Delete old DB snapshots\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.BulkDeleteRdsSnapshotsActionValueFields(),
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
			"change_elasticache_node_type_action_value": {
				Description: "\"ElastiCache: Change node type\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.ChangeElasticacheNodeTypeActionValueFields(),
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
			"delete_nat_gateway_action_value": {
				Description: "\"VPC: Delete NAT Gateway\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.DeleteNatGatewayActionValueFields(),
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
			"no_action_action_value": {
				Description: "\"Other: No Action\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: other.NoActionActionValueFields(),
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
			"ecs_change_service_task_count_action_value": {
				Description: "\"ECS: Change service task count\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.EcsChangeServiceTaskCountActionValueFields(),
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
			"vault_recovery_point_start_copy_job_action_value": {
				Description: "\"Backup: Copy vault recovery point\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.VaultRecoveryPointStartCopyJobActionValueFields(),
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
			"google_compute_start_vm_instances_action_value": {
				Description: "\"Compute Engine: start vm instances\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: gcp.GoogleComputeStartVmInstancesActionValueFields(),
				},
			},
			"google_compute_stop_vm_instances_action_value": {
				Description: "\"Compute Engine: stop vm instances\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: gcp.GoogleComputeStopVmInstancesActionValueFields(),
				},
			},
			"invoke_lambda_function_action_value": {
				Description: "\"Lambda: Invoke lambda function\" action value",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: aws.InvokeLambdaFunctionActionValueFields(),
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
			"amazon_sns_rule_value": {
				Description: "SNS trigger value",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.AmazonSnsRuleValueFields(),
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
			"webhook_rule_value": {
				Description: "HTTP trigger value",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.WebhookRuleValueFields(),
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
				Optional:    true,
				ForceNew:    true,
			},
			"effective_date": {
				Description: "Effective date",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"expiration_date": {
				Description: "Expiration date",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"completed_post_process_id": {
				Description: "Array containing post-process IDs to be executed if the job is successful",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"failed_post_process_id": {
				Description: "Array containing post-process IDs to be executed if the job fails",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceJobCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	job := client.Job{}

	if v, ok := d.GetOk("name"); ok {
		job.Name = v.(string)
	}

	if v, ok := d.GetOk("group_id"); ok {
		job.GroupId = v.(int)
	}

	if v, ok := d.GetOk("aws_account_id"); ok {
		job.AwsAccountId = v.(int)
	}

	if v, ok := d.GetOk("aws_account_ids"); ok && len(v.([]interface{})) > 0 {
		job.AwsAccountIds = utils.ExpandIntList(v.([]interface{}))
	}

	if v, ok := d.GetOk("google_cloud_account_id"); ok {
		job.GoogleCloudAccountId = v.(int)
	}

	if v, ok := d.GetOkExists("for_workflow"); ok {
		fw := v.(bool)
		job.ForWorkflow = &fw
	}

	if v, ok := d.GetOk("rule_type"); ok {
		job.RuleType = v.(string)
	}

	job.RuleValue = buildRuleValue(d, &job)

	if v, ok := d.GetOk("action_type"); ok {
		job.ActionType = v.(string)
	}

	job.ActionValue = buildActionValue(d, &job)

	//lint:ignore SA1019 https://github.com/hashicorp/terraform-plugin-sdk/pull/350#issuecomment-597888969
	if v, ok := d.GetOkExists("allow_runtime_action_values"); ok {
		arav := v.(bool)
		job.AllowRuntimeActionValues = &arav
	}

	if v, ok := d.GetOk("effective_date"); ok {
		job.EffectiveDate = v.(string)
	}

	if v, ok := d.GetOk("expiration_date"); ok {
		job.ExpirationDate = v.(string)
	}

	if v, ok := d.GetOk("completed_post_process_id"); ok && len(v.([]interface{})) > 0 {
		job.CompletedPostProcessId = utils.ExpandIntList(v.([]interface{}))
	}

	if v, ok := d.GetOk("failed_post_process_id"); ok && len(v.([]interface{})) > 0 {
		job.FailedPostProcessId = utils.ExpandIntList(v.([]interface{}))
	}

	j, _, err := c.CreateJob(&job)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(j.Id)

	return resourceJobRead(ctx, d, m)
}

func resourceJobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

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
	d.Set("google_cloud_account_id", job.GoogleCloudAccountId)
	d.Set("for_workflow", job.ForWorkflow)

	d.Set("rule_type", job.RuleType)
	switch job.RuleType {
	case "cron", "schedule", "amazon_sns", "sqs_v2", "webhook":
		ruleValueBlockName := fmt.Sprintf("%s_rule_value", job.RuleType)
		if err := d.Set(ruleValueBlockName, []interface{}{job.RuleValue}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}

	d.Set("action_type", job.ActionType)
	if job.ActionType != "no_action" {
		actionValueBlockName := fmt.Sprintf("%s_action_value", job.ActionType)
		if err := d.Set(actionValueBlockName, []interface{}{buildActionValue(d, job)}); err != nil {
			return append(diags, diag.FromErr(err)...)
		}
	}

	d.Set("allow_runtime_action_values", job.AllowRuntimeActionValues)
	d.Set("effective_date", job.EffectiveDate)
	d.Set("expiration_date", job.ExpirationDate)
	d.Set("completed_post_process_id", utils.FlattenIntList(job.CompletedPostProcessId))
	d.Set("failed_post_process_id", utils.FlattenIntList(job.FailedPostProcessId))

	return diags
}

func resourceJobUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	job, _, err := c.GetJob(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if d.HasChange("name") {
		job.Name = d.Get("name").(string)
	}

	if d.HasChange("group_id") {
		job.GroupId = d.Get("group_id").(int)
	}

	if d.HasChange("aws_account_id") {
		job.AwsAccountId = d.Get("aws_account_id").(int)
	}

	if d.HasChange("aws_account_ids") {
		job.AwsAccountIds = utils.ExpandIntList(d.Get("aws_account_ids").([]interface{}))
	}

	if d.HasChange("google_cloud_account_id") {
		job.GoogleCloudAccountId = d.Get("google_cloud_account_id").(int)
	}

	if d.HasChange("for_workflow") {
		fw := d.Get("for_workflow").(bool)
		job.ForWorkflow = &fw
	}

	ruleValueBlockName := fmt.Sprintf("%s_rule_value", job.RuleType)
	if d.HasChange(ruleValueBlockName) {
		job.RuleValue = buildRuleValue(d, job)
	}

	actionValueBlockName := fmt.Sprintf("%s_action_value", job.ActionType)
	if d.HasChange(actionValueBlockName) {
		job.ActionValue = buildActionValue(d, job)
	}

	if d.HasChange("allow_runtime_action_values") {
		arav := d.Get("allow_runtime_action_values").(bool)
		job.AllowRuntimeActionValues = &arav
	}

	if d.HasChange("effective_date") {
		job.EffectiveDate = d.Get("effective_date").(string)
	}

	if d.HasChange("expiration_date") {
		job.ExpirationDate = d.Get("expiration_date").(string)
	}

	if d.HasChange("completed_post_process_id") {
		job.CompletedPostProcessId = utils.ExpandIntList(d.Get("completed_post_process_id").([]interface{}))
	}

	if d.HasChange("failed_post_process_id") {
		job.FailedPostProcessId = utils.ExpandIntList(d.Get("failed_post_process_id").([]interface{}))
	}

	if _, _, err := c.UpdateJob(job); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return resourceJobRead(ctx, d, m)
}

func resourceJobDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	if _, err := c.DeleteJob(id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId("")

	return nil
}

func buildRuleValue(d *schema.ResourceData, job *client.Job) map[string]interface{} {
	switch job.RuleType {
	case "cron":
		if v, ok := d.GetOk("cron_rule_value"); ok {
			rv := v.([]interface{})[0].(map[string]interface{})

			switch rv["schedule_type"] {
			case "one_time":
				delete(rv, "dates_to_skip")
				delete(rv, "weekly_schedule")
			case "monthly":
				delete(rv, "weekly_schedule")
			case "monthly_day_of_week":
				rv["monthly_schedule"] = rv["monthly_day_of_week_schedule"].([]interface{})[0].(map[string]interface{})
				delete(rv, "monthly_day_of_week_schedule")
				delete(rv, "weekly_schedule")
			}

			if rv["start_timeout_minutes"] == "" {
				delete(rv, "start_timeout_minutes")
			}

			return rv
		}
	case "schedule", "sqs_v2":
		ruleValueBlockName := fmt.Sprintf("%s_rule_value", job.RuleType)
		if v, ok := d.GetOk(ruleValueBlockName); ok {
			return v.([]interface{})[0].(map[string]interface{})
		}
	}

	return nil
}

func buildActionValue(d *schema.ResourceData, job *client.Job) map[string]interface{} {
	blockName := fmt.Sprintf("%s_action_value", job.ActionType)

	var actionValue map[string]interface{}

	// Terraform stateから取得を試みる（Update時など）
	v, ok := d.GetOk(blockName)
	if ok {
		// スライスが空でないことを確認
		actionValueList := v.([]interface{})
		if len(actionValueList) > 0 && actionValueList[0] != nil {
			actionValue = actionValueList[0].(map[string]interface{})
		}
	}

	// stateに値がない場合（Read時など）、APIレスポンスから取得
	if actionValue == nil {
		if job.ActionValue != nil {
			// job.ActionValueをコピーして使用（元のデータを変更しないように）
			actionValue = make(map[string]interface{})
			for k, v := range job.ActionValue {
				actionValue[k] = v
			}
		} else {
			return nil
		}
	}

	schemaFields := resourceJob().Schema[blockName].Elem.(*schema.Resource).Schema
	// job.ActionValueに含まれるキーのうち、スキーマに存在しないものをactionValueから削除
	if job.ActionValue != nil {
		for key := range job.ActionValue {
			if _, exists := schemaFields[key]; !exists {
				delete(actionValue, key)
			}
		}
	}

	switch job.ActionType {
	case "dynamodb_start_backup_job", "ec2_start_backup_job", "s3_start_backup_job", "efs_start_backup_job":
		if actionValue["lifecycle_delete_after_days"] == 0 {
			actionValue["lifecycle_delete_after_days"] = nil
		}

		// additional_tags が存在する場合はリストに変換して、空の場合は空のリストに変換する
		if tagsSet, ok := actionValue["additional_tags"].(*schema.Set); ok && tagsSet.Len() > 0 {
			actionValue["additional_tags"] = tagsSet.List()
		} else {
			actionValue["additional_tags"] = []string{}
		}
	case "vault_recovery_point_start_copy_job":
		if actionValue["lifecycle_delete_after_days"] == 0 {
			actionValue["lifecycle_delete_after_days"] = nil
		}
	default:
		// additional_tags が存在する場合はリストに変換して、空の場合は削除する
		if tags, ok := actionValue["additional_tags"]; ok {
			if tagsSet, ok := tags.(*schema.Set); ok && tagsSet.Len() > 0 {
				actionValue["additional_tags"] = tagsSet.List()
			} else {
				delete(actionValue, "additional_tags")
			}
		}
	}

	return actionValue
}
