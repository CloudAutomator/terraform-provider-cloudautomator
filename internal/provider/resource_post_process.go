package cloudautomator

import (
	"context"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"
	schemes "terraform-provider-cloudautomator/internal/schemes/post_process"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePostProcess() *schema.Resource {
	return &schema.Resource{
		Description:   "Manage Cloud Automator post-process resources",
		CreateContext: resourcePostProcessCreate,
		ReadContext:   resourcePostProcessRead,
		UpdateContext: resourcePostProcessUpdate,
		DeleteContext: resourcePostProcessDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Post-process ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Post-process name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"service": {
				Description: "Post-process service name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"group_id": {
				Description: "Group ID",
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
			},
			"shared_by_group": {
				Description: "Whether shared by groups",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"email_parameters": {
				Description: "\"email\" parameter value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.EmailParametersFields(),
				},
			},
			"slack_parameters": {
				Description: "slack parameter value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.SlackParametersFields(),
				},
			},
			"sqs_parameters": {
				Description: "SQS parameter value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.SqsParametersFields(),
				},
			},
			"webhook_parameters": {
				Description: "webhook parameter value",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: schemes.WebhookParametersFields(),
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePostProcessCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	postProcess := client.PostProcess{}

	if v, ok := d.GetOk("name"); ok {
		postProcess.Name = v.(string)
	}

	if v, ok := d.GetOk("service"); ok {
		postProcess.Service = v.(string)
	}

	if v, ok := d.GetOk("group_id"); ok {
		postProcess.GroupId = v.(int)
	}

	//lint:ignore SA1019 https://github.com/hashicorp/terraform-plugin-sdk/pull/350#issuecomment-597888969
	if v, ok := d.GetOkExists("shared_by_group"); ok {
		sbg := v.(bool)
		postProcess.SharedByGroup = &sbg
	}

	postProcess.Parameters = buildParameters(d, &postProcess)

	p, _, err := c.CreatePostProcess(&postProcess)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(p.Id)

	return resourcePostProcessRead(ctx, d, m)
}

func resourcePostProcessRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	postProcess, _, err := c.GetPostProcess(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("id", postProcess.Id)
	d.Set("name", postProcess.Name)
	d.Set("service", postProcess.Service)
	d.Set("group_id", postProcess.GroupId)
	d.Set("shared_by_group", postProcess.SharedByGroup)

	parameterBlockName := fmt.Sprintf("%s_parameters", postProcess.Service)

	if err := d.Set(parameterBlockName, []interface{}{postProcess.Parameters}); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return diags
}

func resourcePostProcessUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	postProcess, _, err := c.GetPostProcess(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if d.HasChange("name") {
		postProcess.Name = d.Get("name").(string)
	}

	if d.HasChange("aws_account_id") {
		postProcess.Service = d.Get("service").(string)
	}

	if d.HasChange("group_id") {
		postProcess.GroupId = d.Get("group_id").(int)
	}

	if d.HasChange("shared_by_group") {
		sbg := d.Get("shared_by_group").(bool)
		postProcess.SharedByGroup = &sbg
	}

	parameterBlockName := fmt.Sprintf("%s_parameters", postProcess.Parameters)
	if d.HasChange(parameterBlockName) {
		postProcess.Parameters = buildParameters(d, postProcess)
	}

	if _, _, err := c.UpdatePostProcess(postProcess); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return resourcePostProcessRead(ctx, d, m)
}

func resourcePostProcessDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	if _, err := c.DeletePostProcess(id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId("")

	return nil
}

func buildParameters(d *schema.ResourceData, postProcess *client.PostProcess) map[string]interface{} {
	parameterBlockName := fmt.Sprintf("%s_parameters", postProcess.Service)

	if v, ok := d.GetOk(parameterBlockName); ok {
		return v.([]interface{})[0].(map[string]interface{})
	}

	return nil
}
