package cloudautomator

import (
	"context"
	"errors"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ApiEndpoint = "CLOUD_AUTOMATOR_API_ENDPOINT"
const ApiKeyEnvName = "CLOUD_AUTOMATOR_API_KEY"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(ApiEndpoint, client.ApiEndpoint),
				Description: fmt.Sprintf("Cloud Automator API Endpoint. This can also be set via the %s environment variable.", ApiEndpoint),
			},
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
	var c *client.Client
	var diags diag.Diagnostics

	apiKey := d.Get("api_key").(string)

	if apiKey == "" {
		return nil, diag.FromErr(errors.New("api_key must be set"))
	}

	apiEndPoint := d.Get("api_endpoint").(string)
	clientOption := client.WithAPIEndpoint(apiEndPoint)

	c, _ = client.New(apiKey, clientOption)
	return c, diags
}
