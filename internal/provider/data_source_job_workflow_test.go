package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorDataSourceJobWorkflow_basic(t *testing.T) {
	dataSourceName := "cloudautomator_job_workflow.test"
	jobWorkflowName := fmt.Sprintf("tf-testacc-job-workflow-%s", utils.RandomString(12))
	groupId := acctest.TestGroupId()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAutomatorDataSourceJobWorkflow_basic(jobWorkflowName, groupId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", jobWorkflowName),
					resource.TestCheckResourceAttr(dataSourceName, "group_id", groupId),
					resource.TestCheckResourceAttrPair(dataSourceName, "first_job_id", "cloudautomator_job.first_job", "id"),
					resource.TestCheckResourceAttr(dataSourceName, "following_job_ids.#", "1"),
					resource.TestCheckResourceAttrPair(dataSourceName, "following_job_ids.0", "cloudautomator_job.following_job", "id"),
				),
			},
		},
	})
}

func testAccCloudAutomatorDataSourceJobWorkflow_basic(jobWorkflowName string, groupId string) string {
	return fmt.Sprintf(`
resource "cloudautomator_job" "first_job" {
  name      = "first_job"
  group_id  = %s

  rule_type = "webhook"

  for_workflow = true

  action_type = "no_action"
}

resource "cloudautomator_job" "following_job" {
	name     = "following_job"
	group_id = %s

	for_workflow = true

	rule_type = "no_rule"

	action_type = "delay"
	delay_action_value {
	  delay_minutes = 1
	}
}

resource "cloudautomator_job_workflow" "test" {
	name              = "%s"
	group_id          = %s
	first_job_id      = cloudautomator_job.first_job.id
	following_job_ids = [cloudautomator_job.following_job.id]
}

data "cloudautomator_job_workflow" "test" {
	id = cloudautomator_job_workflow.test.id
}
`, groupId, groupId, jobWorkflowName, groupId)
}
