package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/client"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudAutomatorJobWorkflow(t *testing.T) {
	cases := []struct {
		name            string
		jobWorkflowName string
		configFunc      func(string, string) string
		checks          []resource.TestCheckFunc
	}{
		{
			name:            "BasicJobWorkflow",
			jobWorkflowName: fmt.Sprintf("tf-testacc-job-workflow-%s", utils.RandomString(12)),
			configFunc: func(jobWorkflowName string, groupId string) string {
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
				`, groupId, groupId, jobWorkflowName, groupId)
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job_workflow.test", "group_id", acctest.TestGroupId()),
				resource.TestCheckResourceAttrPair("cloudautomator_job_workflow.test", "first_job_id", "cloudautomator_job.first_job", "id"),
				resource.TestCheckResourceAttr("cloudautomator_job_workflow.test", "following_job_ids.#", "1"),
				resource.TestCheckResourceAttrPair("cloudautomator_job_workflow.test", "following_job_ids.0", "cloudautomator_job.following_job", "id"),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			groupId := acctest.TestGroupId()

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testAccCheckCloudAutomatorJobWorkflowDestroy,
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(tc.jobWorkflowName, groupId),
						Check: resource.ComposeTestCheckFunc(
							append([]resource.TestCheckFunc{
								testAccCheckCloudAutomatorJobWorkflowExists(testAccProviders["cloudautomator"], "cloudautomator_job_workflow.test"),
								resource.TestCheckResourceAttr("cloudautomator_job_workflow.test", "name", tc.jobWorkflowName),
							}, tc.checks...)...,
						),
					},
				},
			})
		})
	}
}

func testAccCheckCloudAutomatorJobWorkflowExists(_ *schema.Provider, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*client.Client)

		if err := cloudautomatorJobWorkflowExistsHelper(s, c, n); err != nil {
			return err
		}

		return nil
	}
}

func cloudautomatorJobWorkflowExistsHelper(s *terraform.State, c *client.Client, name string) error {
	id := s.RootModule().Resources[name].Primary.ID
	if _, _, err := c.GetJobWorkflow(id); err != nil {
		return fmt.Errorf("received an error retrieving job workflow %s", err)
	}

	return nil
}

func testAccCheckCloudAutomatorJobWorkflowDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudautomator_job_workflow" {
			continue
		}

		jobWorkflowId := rs.Primary.ID

		_, res, err := c.GetJobWorkflow(jobWorkflowId)
		if err != nil {
			if res.StatusCode == 404 {
				continue
			}

			return fmt.Errorf("received an error retrieving job workflow %s", err)
		}

		return fmt.Errorf("job workflow exists.")
	}

	return nil
}
