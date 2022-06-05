package cloudautomator

import (
	"context"
	"errors"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ApiKeyEnvName = "CA_API_KEY"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(ApiKeyEnvName, nil),
				Description: fmt.Sprintf("Cloud Automator API key. This can also be set via the %s environment variable.", ApiKeyEnvName),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudautomator_job":          resourceJob(),
			"cloudautomator_post_process": resourcePostProcess(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudautomator_aws_account":  dataSourceAwsAccount(),
			"cloudautomator_job":          dataSourceJob(),
			"cloudautomator_post_process": dataSourcePostProcess(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)

	var diags diag.Diagnostics

	if apiKey == "" {
		return nil, diag.FromErr(errors.New("api_key must be set"))
	}

	client, _ := client.New(&apiKey)

	return client, diags
}
