package cloudautomator

import (
	"context"
	"terraform-provider-cloudautomator/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAwsAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAwsAccountRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "AWSアカウントID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"group_id": {
				Description: "グループID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "AWSアカウント名",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceAwsAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	var diags diag.Diagnostics
	id := d.Get("id").(string)
	group_id := d.Get("group_id").(string)

	awsAccount, _, err := c.GetAwsAccount(group_id, id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.Set("id", awsAccount.Id)
	d.Set("name", awsAccount.Name)

	d.SetId(awsAccount.Id)

	return diags
}
