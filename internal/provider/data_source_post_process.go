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
				Description: "Post-process ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Post-process name",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service": {
				Description: "Post-process service name",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"group_id": {
				Description: "Group ID",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"shared_by_group": {
				Description: "Whether shared by groups",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"email_parameters": {
				Description: "\"email\" parameter value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.EmailParametersFields(),
				},
			},
			"slack_parameters": {
				Description: "slack parameter value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.SlackParametersFields(),
				},
			},
			"sqs_parameters": {
				Description: "SQS parameter value",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: schemes.SqsParametersFields(),
				},
			},
			"webhook_parameters": {
				Description: "webhook parameter value",
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
