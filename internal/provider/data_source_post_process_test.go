package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorDataSourcePostProcess_basic(t *testing.T) {
	dataSourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAutomatorDataSourcePostProcess_basic(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						dataSourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						dataSourceName, "service", "email"),
					resource.TestCheckResourceAttr(
						dataSourceName, "group_id", acctest.TestGroupId()),
					resource.TestCheckResourceAttr(
						dataSourceName, "shared_by_group", "false"),
					resource.TestCheckResourceAttr(
						dataSourceName, "email_parameters.0.email_recipient", "test@example.com"),
				),
			},
		},
	})
}

func testAccCloudAutomatorDataSourcePostProcess_basic(rName string) string {
	return fmt.Sprintf(`
resource "cloudautomator_post_process" "test" {
	name = "%s"
	group_id = "%s"
	service = "email"

	email_parameters {
		email_recipient = "test@example.com"
	}
}

data "cloudautomator_post_process" "test" {
	id = cloudautomator_post_process.test.id
}`, rName, acctest.TestGroupId())
}
