package cloudautomator

import (
	"context"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"
	schemes "terraform-provider-cloudautomator/internal/schemes/job"
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
			"authorize_security_group_ingress_action_value": {
				Description: "\"EC2: Authorize security group ingress\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.AuthorizeSecurityGroupIngressyActionValueFields(),
				},
			},
			"change_rds_cluster_instance_class_action_value": {
				Description: "\"RDS(Aurora): Change DB instance class\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.ChangeRdsClusterInstanceClassActionValueFields(),
				},
			},
			"change_rds_instance_class_action_value": {
				Description: "\"RDS: Change DB instance class\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.ChangeRdsInstanceClassActionValueFields(),
				},
			},
			"change_instance_type_action_value": {
				Description: "\"EC2: Change instance type\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.ChangeInstanceTypeActionValueFields(),
				},
			},
			"copy_ebs_snapshot_action_value": {
				Description: "\"EC2: Copy EBS snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CopyEbsSnapshotActionValueFields(),
				},
			},
			"copy_image_action_value": {
				Description: "\"EC2: Copy AMI\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CopyImageActionValueFields(),
				},
			},
			"copy_rds_snapshot_action_value": {
				Description: "\"RDS: Copy DB snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CopyRdsSnapshotActionValueFields(),
				},
			},
			"create_ebs_snapshot_action_value": {
				Description: "\"EC2: Create EBS snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CreateEbsSnapshotActionValueFields(),
				},
			},
			"create_image_action_value": {
				Description: "\"EC2: Create AMI\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CreateImageActionValueFields(),
				},
			},
			"create_rds_cluster_snapshot_action_value": {
				Description: "\"RDS(Aurora): Create DB cluster snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CreateRdsClusterSnapshotActionValueFields(),
				},
			},
			"create_rds_snapshot_action_value": {
				Description: "\"RDS: Create DB snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CreateRdsSnapshotActionValueFields(),
				},
			},
			"create_redshift_snapshot_action_value": {
				Description: "\"Redshift: Create cluster snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.CreateRedshiftSnapshotActionValueFields(),
				},
			},
			"delay_action_value": {
				Description: "\"Other: Delay\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.DelayActionValueFields(),
				},
			},
			"delete_cluster_action_value": {
				Description: "\"Redshift: Delete cluster\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.DeleteClusterActionValueFields(),
				},
			},
			"delete_rds_cluster_action_value": {
				Description: "\"RDS(Aurora): Delete DB cluster\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.DeleteRdsClusterActionValueFields(),
				},
			},
			"delete_rds_instance_action_value": {
				Description: "\"RDS: Delete DB instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.DeleteRdsInstanceActionValueFields(),
				},
			},
			"deregister_instances_action_value": {
				Description: "\"ELB(CLB): De-register EC2 instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.DeregisterInstancesActionValueFields(),
				},
			},
			"deregister_target_instances_action_value": {
				Description: "\"ELB(ALB/NLB): Deregister EC2 instances from target group\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.DeregisterTargetInstancesActionValueFields(),
				},
			},
			"reboot_rds_instances_action_value": {
				Description: "\"RDS: Reboot DB instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RebootRdsInstancesActionValueFields(),
				},
			},
			"reboot_workspaces_action_value": {
				Description: "\"WorkSpaces: Reboot WorkSpace\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RebootWorkspacesActionValueFields(),
				},
			},
			"register_instances_action_value": {
				Description: "\"ELB(CLB): Register EC2 instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RegisterInstancesActionValueFields(),
				},
			},
			"register_target_instances_action_value": {
				Description: "\"ELB(ALB/NLB): Register EC2 instances to target group\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RegisterTargetInstancesActionValueFields(),
				},
			},
			"restore_from_cluster_snapshot_action_value": {
				Description: "\"Redshift: Restore from snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RestoreFromClusterSnapshotActionValueFields(),
				},
			},
			"restore_rds_cluster_action_value": {
				Description: "\"RDS(Aurora): Restore DB cluster from DB cluster snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RestoreRdsClusterActionValueFields(),
				},
			},
			"restore_rds_instance_action_value": {
				Description: "\"RDS: Restore from DB snapshot\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RestoreRdsInstanceActionValueFields(),
				},
			},
			"revoke_security_group_ingress_action_value": {
				Description: "\"EC2: Revoke security group ingress\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.RevokeSecurityGroupIngressActionValueFields(),
				},
			},
			"send_command_action_value": {
				Description: "\"EC2: Send command on instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.SendCommandActionValueFields(),
				},
			},
			"start_instances_action_value": {
				Description: "\"EC2: Start instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StartInstancesActionValueFields(),
				},
			},
			"start_rds_clusters_action_value": {
				Description: "\"RDS(Aurora): Start DB cluster\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StartRdsClustersActionValueFields(),
				},
			},
			"start_rds_instances_action_value": {
				Description: "\"RDS: Start DB instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StartRdsInstancesActionValueFields(),
				},
			},
			"stop_instances_action_value": {
				Description: "\"EC2: Stop instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StopInstancesActionValueFields(),
				},
			},
			"stop_rds_clusters_action_value": {
				Description: "\"RDS(Aurora): Stop DB cluster\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StopRdsClustersActionValueFields(),
				},
			},
			"stop_rds_instances_action_value": {
				Description: "\"RDS: Stop DB instance\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StopRdsInstancesActionValueFields(),
				},
			},
			"start_workspaces_action_value": {
				Description: "\"WorkSpaces: Start WorkSpace\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.StartWorkspacesActionValueFields(),
				},
			},
			"terminate_workspaces_action_value": {
				Description: "\"WorkSpaces: Remove WorkSpace\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.TerminateWorkspacesActionValueFields(),
				},
			},
			"update_record_set_action_value": {
				Description: "\"Route 53: Update Resource Record Set\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.UpdateRecordSetActionValueFields(),
				},
			},
			"windows_update_action_value": {
				Description: "\"EC2: Windows Update to instance (Old version)\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.WindowsUpdateActionValueFields(),
				},
			},
			"windows_update_v2_action_value": {
				Description: "\"EC2: Windows Update to instance (New version)\" action value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.WindowsUpdateV2ActionValueFields(),
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
