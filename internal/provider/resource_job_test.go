package cloudautomator

import (
	"fmt"

	"terraform-provider-cloudautomator/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccCheckCloudAutomatorJobExists(_ *schema.Provider, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*client.Client)

		if err := cloudautomatorJobExistsHelper(s, c, n); err != nil {
			return err
		}

		return nil
	}
}

func cloudautomatorJobExistsHelper(s *terraform.State, c *client.Client, name string) error {
	id := s.RootModule().Resources[name].Primary.ID
	if _, _, err := c.GetJob(id); err != nil {
		return fmt.Errorf("received an error retrieving job %s", err)
	}

	return nil
}

func testAccCheckCloudAutomatorJobDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudautomator_job" {
			continue
		}

		jobId := rs.Primary.ID

		_, res, err := c.GetJob(jobId)
		if err != nil {
			if res.StatusCode == 404 {
				continue
			}

			return fmt.Errorf("received an error retrieving job %s", err)
		}

		return fmt.Errorf("job exists.")
	}

	return nil
}
