package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_DelayAction(t *testing.T) {
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				rule_type = "webhook"

				action_type = "delay"
				delay_action_value {
					delay_minutes = 30
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, jobName, acctest.TestGroupId(), acctest.TestPostProcessId(), acctest.TestPostProcessId()),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], "cloudautomator_job.test"),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "name", jobName),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delay"),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "delay_action_value.0.delay_minutes", "30"),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "completed_post_process_id.0", acctest.TestPostProcessId()),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "failed_post_process_id.0", acctest.TestPostProcessId()),
				),
			},
		},
	})
}

func TestAccCloudAutomatorJob_NoActionAction(t *testing.T) {
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"

				for_workflow = true

				rule_type = "webhook"

				action_type = "no_action"
			}`, jobName, acctest.TestGroupId()),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], "cloudautomator_job.test"),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "name", jobName),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "no_action"),
					resource.TestCheckResourceAttr("cloudautomator_job.test", "for_workflow", "true"),
				),
			},
		},
	})
}
