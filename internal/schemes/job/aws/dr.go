package schemes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DescribeMetadataActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dr_configuration_id": {
			Description: "DR Configuration ID",
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func DisasterRecoveryActionValueFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"describe_metadata_job_id": {
			Description: "Describe Metadata Job ID",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"trace_status": {
			Description: "Whether to Verify completion status of the resource",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
