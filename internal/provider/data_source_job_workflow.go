package cloudautomator

import (
	"context"
	"terraform-provider-cloudautomator/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJobWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJobWorkflowRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Job Workflow ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"first_job_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"following_job_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceJobWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Get("id").(string)

	jobWorkflow, _, err := c.GetJobWorkflow(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("id", jobWorkflow.Id)
	d.Set("name", jobWorkflow.Name)
	d.Set("group_id", jobWorkflow.GroupId)
	d.Set("first_job_id", jobWorkflow.FirstJobId)
	d.Set("following_job_ids", jobWorkflow.FollowingJobIds)

	d.SetId(jobWorkflow.Id)

	return diags
}
