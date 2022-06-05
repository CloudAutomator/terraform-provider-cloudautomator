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

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Description:   "Cloud Automatorのジョブリソースを管理します。",
		CreateContext: resourceJobCreate,
		ReadContext:   resourceJobRead,
		UpdateContext: resourceJobUpdate,
		DeleteContext: resourceJobDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "ジョブID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "ジョブ名",
				Type:        schema.TypeString,
				Required:    true,
			},
			"group_id": {
				Description: "グループID",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"aws_account_id": {
				Description: "AWSアカウントID",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"rule_type": {
				Description: "トリガーのタイプ",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"cron_rule_value": {
				Description: "タイマートリガーの設定値",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.CronRuleValueFields(),
				},
			},
			"schedule_rule_value": {
				Description: "スケジュールトリガーの設定値",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.ScheduleRuleValueFields(),
				},
			},
			"sqs_v2_rule_value": {
				Description: "SQSトリガーの設定値",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.SqsV2RuleValueFields(),
				},
			},
			"action_type": {
				Description: "アクションのタイプ",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"authorize_security_group_ingress_action_value": {
				Description: "「EC2: セキュリティグループにインバウンドルールを追加」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.AuthorizeSecurityGroupIngressyActionValueFields(),
				},
			},
			"change_rds_cluster_instance_class_action_value": {
				Description: "「RDS(Aurora): DBインスタンスクラスを変更」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.ChangeRdsClusterInstanceClassActionValueFields(),
				},
			},
			"change_rds_instance_class_action_value": {
				Description: "「RDS: DBインスタンスクラスを変更」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.ChangeRdsInstanceClassActionValueFields(),
				},
			},
			"change_instance_type_action_value": {
				Description: "「EC2: インスタンスタイプを変更」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.ChangeInstanceTypeActionValueFields(),
				},
			},
			"copy_ebs_snapshot_action_value": {
				Description: "「EC2: EBSスナップショットをリージョン間でコピー」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CopyEbsSnapshotActionValueFields(),
				},
			},
			"copy_image_action_value": {
				Description: "「EC2: AMIをリージョン間でコピー」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CopyImageActionValueFields(),
				},
			},
			"copy_rds_snapshot_action_value": {
				Description: "「RDS: DBスナップショットをリージョン間でコピー」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CopyRdsSnapshotActionValueFields(),
				},
			},
			"create_ebs_snapshot_action_value": {
				Description: "「EC2: EBSスナップショットを作成」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CreateEbsSnapshotActionValueFields(),
				},
			},
			"create_image_action_value": {
				Description: "「EC2: AMIを作成」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CreateImageActionValueFields(),
				},
			},
			"create_rds_cluster_snapshot_action_value": {
				Description: "「RDS(Aurora): DBクラスタースナップショットを作成」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CreateRdsClusterSnapshotActionValueFields(),
				},
			},
			"create_rds_snapshot_action_value": {
				Description: "「RDS: DBスナップショットを作成」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CreateRdsSnapshotActionValueFields(),
				},
			},
			"create_redshift_snapshot_action_value": {
				Description: "「Redshift: クラスタースナップショットを作成」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.CreateRedshiftSnapshotActionValueFields(),
				},
			},
			"delay_action_value": {
				Description: "「Other: 指定時間待機」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.DelayActionValueFields(),
				},
			},
			"delete_cluster_action_value": {
				Description: "「Redshift: クラスターを削除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.DeleteClusterActionValueFields(),
				},
			},
			"delete_rds_cluster_action_value": {
				Description: "「RDS(Aurora): DBクラスターを削除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.DeleteRdsClusterActionValueFields(),
				},
			},
			"delete_rds_instance_action_value": {
				Description: "「RDS: DBインスタンスを削除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.DeleteRdsInstanceActionValueFields(),
				},
			},
			"deregister_instances_action_value": {
				Description: "「ELB(CLB): EC2インスタンスを登録解除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.DeregisterInstancesActionValueFields(),
				},
			},
			"deregister_target_instances_action_value": {
				Description: "「ELB(ALB/NLB): ターゲットグループからEC2インスタンスを登録解除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.DeregisterTargetInstancesActionValueFields(),
				},
			},
			"reboot_rds_instances_action_value": {
				Description: "「RDS: DBインスタンスを再起動」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RebootRdsInstancesActionValueFields(),
				},
			},
			"reboot_workspaces_action_value": {
				Description: "「WorkSpaces: WorkSpaceを再起動」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RebootWorkspacesActionValueFields(),
				},
			},
			"register_instances_action_value": {
				Description: "「ELB(CLB): EC2インスタンスを登録」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RegisterInstancesActionValueFields(),
				},
			},
			"register_target_instances_action_value": {
				Description: "「ELB(ALB/NLB): ターゲットグループにEC2インスタンスを登録」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RegisterTargetInstancesActionValueFields(),
				},
			},
			"restore_from_cluster_snapshot_action_value": {
				Description: "「Redshift: スナップショットからリストア」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RestoreFromClusterSnapshotActionValueFields(),
				},
			},
			"restore_rds_cluster_action_value": {
				Description: "「RDS(Aurora): DBクラスタースナップショットからリストア」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RestoreRdsClusterActionValueFields(),
				},
			},
			"restore_rds_instance_action_value": {
				Description: "「RDS: DBスナップショットからリストア」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RestoreRdsInstanceActionValueFields(),
				},
			},
			"revoke_security_group_ingress_action_value": {
				Description: "「EC2: セキュリティグループからインバウンドルールを削除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.RevokeSecurityGroupIngressActionValueFields(),
				},
			},
			"send_command_action_value": {
				Description: "「EC2: インスタンスでコマンドを実行」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.SendCommandActionValueFields(),
				},
			},
			"start_instances_action_value": {
				Description: "「EC2: インスタンスを起動」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StartInstancesActionValueFields(),
				},
			},
			"start_rds_clusters_action_value": {
				Description: "「RDS(Aurora): DBクラスターを起動」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StartRdsClustersActionValueFields(),
				},
			},
			"start_rds_instances_action_value": {
				Description: "「RDS: DBインスタンスを起動」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StartRdsInstancesActionValueFields(),
				},
			},
			"stop_instances_action_value": {
				Description: "「EC2: インスタンスを停止」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StopInstancesActionValueFields(),
				},
			},
			"stop_rds_clusters_action_value": {
				Description: "「RDS(Aurora): DBクラスターを停止」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StopRdsClustersActionValueFields(),
				},
			},
			"stop_rds_instances_action_value": {
				Description: "「RDS: DBインスタンスを停止」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StopRdsInstancesActionValueFields(),
				},
			},
			"start_workspaces_action_value": {
				Description: "「WorkSpaces: WorkSpaceを起動」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.StartWorkspacesActionValueFields(),
				},
			},
			"terminate_workspaces_action_value": {
				Description: "「WorkSpaces: WorkSpaceを削除」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.TerminateWorkspacesActionValueFields(),
				},
			},
			"update_record_set_action_value": {
				Description: "「Route 53: リソースレコードセットを更新」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.UpdateRecordSetActionValueFields(),
				},
			},
			"windows_update_action_value": {
				Description: "「EC2: インスタンスをWindows Update」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.WindowsUpdateActionValueFields(),
				},
			},
			"windows_update_v2_action_value": {
				Description: "「EC2: インスタンスをWindows Update (新バージョン)」アクションの設定値",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: schemes.WindowsUpdateV2ActionValueFields(),
				},
			},
			"allow_runtime_action_values": {
				Description: "アクションの設定値を実行時に指定するかどうか",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"effective_date": {
				Description: "ジョブの有効期間の開始日",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"expiration_date": {
				Description: "ジョブの有効期間の終了日",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"completed_post_process_id": {
				Description: "ジョブが成功した場合に実行する後処理IDが含まれる配列",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"failed_post_process_id": {
				Description: "ジョブが失敗した場合に実行する後処理IDが含まれる配列",
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
	switch job.ActionType {
	case "create_image":
		if v, ok := d.GetOk("create_image_action_value"); ok {
			actionValue := v.([]interface{})[0].(map[string]interface{})
			actionValue["additional_tags"] = actionValue["additional_tags"].(*schema.Set).List()

			return actionValue
		}
	default:
		blockName := fmt.Sprintf("%s_action_value", job.ActionType)
		if v, ok := d.GetOk(blockName); ok {
			return v.([]interface{})[0].(map[string]interface{})
		}
	}

	return nil
}
