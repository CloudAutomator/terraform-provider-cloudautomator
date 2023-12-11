package cloudautomator

import (
	"context"

	"terraform-provider-cloudautomator/internal/client"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceJobWorkflow() *schema.Resource {
	return &schema.Resource{
		Description:   "Manage Cloud Automator job workflow resources",
		CreateContext: resourceJobWorkflowCreate,
		ReadContext:   resourceJobWorkflowRead,
		UpdateContext: resourceJobWorkflowUpdate,
		DeleteContext: resourceJobWorkflowDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Job Workflow ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"first_job_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"following_job_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceJobWorkflowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	jobWorkflow := client.JobWorkflow{}

	if v, ok := d.GetOk("name"); ok {
		jobWorkflow.Name = v.(string)
	}

	if v, ok := d.GetOkExists("active"); ok {
		activeValue := v.(bool)
		jobWorkflow.Active = &activeValue
	}

	if v, ok := d.GetOk("group_id"); ok {
		jobWorkflow.GroupId = v.(int)
	}

	if v, ok := d.GetOk("first_job_id"); ok {
		jobWorkflow.FirstJobId = v.(int)
	}

	if v, ok := d.GetOk("following_job_ids"); ok && len(v.([]interface{})) > 0 {
		jobWorkflow.FollowingJobIds = utils.ExpandIntList(v.([]interface{}))
	}

	jw, _, err := c.CreateJobWorkflow(&jobWorkflow)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(jw.Id)

	return resourceJobWorkflowRead(ctx, d, m)
}

func resourceJobWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	jobWorkflow, _, err := c.GetJobWorkflow(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("id", jobWorkflow.Id)
	d.Set("active", jobWorkflow.Active)
	d.Set("name", jobWorkflow.Name)
	d.Set("group_id", jobWorkflow.GroupId)

	return diags
}

func resourceJobWorkflowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	jobWorkflow, _, err := c.GetJobWorkflow(id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if d.HasChange("name") {
		jobWorkflow.Name = d.Get("name").(string)
	}

	if v, ok := d.GetOk("active"); ok {
		activeValue := v.(bool)
		jobWorkflow.Active = &activeValue
	}

	if d.HasChange("group_id") {
		jobWorkflow.GroupId = d.Get("group_id").(int)
	}

	if d.HasChange("first_job_id") {
		jobWorkflow.FirstJobId = d.Get("first_job_id").(int)
	}

	if d.HasChange("following_job_ids") {
		jobWorkflow.FollowingJobIds = utils.ExpandIntList(d.Get("following_job_ids").([]interface{}))
	}

	if _, _, err := c.UpdateJobWorkflow(jobWorkflow); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return resourceJobWorkflowRead(ctx, d, m)
}

func resourceJobWorkflowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Id()

	if _, err := c.DeleteJobWorkflow(id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId("")

	return nil
}
