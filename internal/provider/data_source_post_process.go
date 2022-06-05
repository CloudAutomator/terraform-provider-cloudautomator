package cloudautomator

import (
	"context"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"
	schemes "terraform-provider-cloudautomator/internal/schemes/post_process"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePostProcess() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePostProcessRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "後処理ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "後処理名",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service": {
				Description: "サービス名",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"group_id": {
				Description: "グループID",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"shared_by_group": {
				Description: "共通後処理にするかどうか",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"email_parameters": {
				Description: "メールの設定値",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.EmailParametersFields(),
				},
			},
			"slack_parameters": {
				Description: "メールの設定値",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.SlackParametersFields(),
				},
			},
			"sqs_parameters": {
				Description: "メールの設定値",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.SqsParametersFields(),
				},
			},
			"webhook_parameters": {
				Description: "メールの設定値",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.WebhookParametersFields(),
				},
			},
		},
	}
}

func dataSourcePostProcessRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Get("id").(string)

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

	d.SetId(postProcess.Id)

	return diags
}
