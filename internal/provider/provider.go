package cloudautomator

import (
	"context"
	"errors"
	"fmt"
	"terraform-provider-cloudautomator/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const apiEndpointEnvName = "CLOUD_AUTOMATOR_API_ENDPOINT"
const apiKeyEnvName = "CLOUD_AUTOMATOR_API_KEY"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiEndpointEnvName, nil),
				Description: fmt.Sprintf("Cloud Automator API Endpoint. This can also be set via the %s environment variable.", apiEndpointEnvName),
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(apiKeyEnvName, nil),
				Description: fmt.Sprintf("Cloud Automator API key. This can also be set via the %s environment variable.", apiKeyEnvName),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudautomator_job":          resourceJob(),
			"cloudautomator_job_workflow": resourceJobWorkflow(),
			"cloudautomator_post_process": resourcePostProcess(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudautomator_aws_account":  dataSourceAwsAccount(),
			"cloudautomator_job":          dataSourceJob(),
			"cloudautomator_job_workflow": dataSourceJobWorkflow(),
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

	var clientOptions []client.ClientOptions
	if apiEndPoint != "" {
		clientOptions = append(clientOptions, client.WithAPIEndpoint(apiEndPoint))
	}

	c, err := client.New(apiKey, clientOptions...)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("failed to create Cloud Automator client: %w", err))
	}

	return c, diags
}
