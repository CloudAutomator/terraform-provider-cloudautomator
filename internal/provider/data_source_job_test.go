package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorDataSourceJob_basic(t *testing.T) {
	dataSourceName := "cloudautomator_job.test"
	jobName := fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAutomatorDataSourceJob_basic(jobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						dataSourceName, "name", jobName),
					resource.TestCheckResourceAttr(
						dataSourceName, "rule_type", "cron"),
					resource.TestCheckResourceAttr(
						dataSourceName, "cron_rule_value.0.hour", "3"),
					resource.TestCheckResourceAttr(
						dataSourceName, "cron_rule_value.0.minutes", "30"),
					resource.TestCheckResourceAttr(
						dataSourceName, "cron_rule_value.0.schedule_type", "one_time"),
					resource.TestCheckResourceAttr(
						dataSourceName, "cron_rule_value.0.one_time_schedule", "2099/01/01"),
					resource.TestCheckResourceAttr(
						dataSourceName, "cron_rule_value.0.time_zone", "Tokyo"),
				),
			},
		},
	})
}

func testAccCloudAutomatorDataSourceJob_basic(rName string) string {
	return fmt.Sprintf(`
resource "cloudautomator_job" "test" {
	name = "%s"
	group_id = "%s"
	aws_account_id = "%s"

	rule_type = "cron"
	cron_rule_value {
		hour = "3"
		minutes = "30"
		schedule_type = "one_time"
		one_time_schedule = "2099/01/01"
		time_zone = "Tokyo"
	}

	action_type = "delay"
	delay_action_value {
		delay_minutes = 1
	}
}

data "cloudautomator_job" "test" {
	id = cloudautomator_job.test.id
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId())
}
