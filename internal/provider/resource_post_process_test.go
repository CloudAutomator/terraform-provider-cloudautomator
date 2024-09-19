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

func TestAccCloudAutomatorPostProcess(t *testing.T) {
	cases := []struct {
		name            string
		postProcessName string
		configFunc      func(string) string
		checks          []resource.TestCheckFunc
	}{
		{
			name:            "Email",
			postProcessName: fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
				resource "cloudautomator_post_process" "test" {
					name = "%s"
					group_id = "%s"
					service = "email"

					email_parameters {
						email_recipient = "test@example.com"
					}
				}`, resourceName, acctest.TestGroupId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "service", "email"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "email_parameters.0.email_recipient", "test@example.com"),
			},
		},
		{
			name:            "Slack",
			postProcessName: fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
				resource "cloudautomator_post_process" "test" {
					name = "%s"
					group_id = "%s"
					service = "slack"

					slack_parameters {
						slack_channel_name = "ca-test-info-notification"
						slack_language = "ja"
						slack_time_zone = "Tokyo"
					}
				}`, resourceName, acctest.TestGroupId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "service", "slack"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "slack_parameters.0.slack_channel_name", "ca-test-info-notification"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "slack_parameters.0.slack_language", "ja"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "slack_parameters.0.slack_time_zone", "Tokyo"),
			},
		},
		{
			name:            "Sqs",
			postProcessName: fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
				resource "cloudautomator_post_process" "test" {
					name = "%s"
					group_id = "%s"
					service = "sqs"

					sqs_parameters {
						sqs_aws_account_id = "%s"
						sqs_queue = "test-queue"
						sqs_region = "ap-northeast-1"
					}
				}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "service", "sqs"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "sqs_parameters.0.sqs_queue", "test-queue"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "sqs_parameters.0.sqs_region", "ap-northeast-1"),
			},
		},
		{
			name:            "Webhook",
			postProcessName: fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
				resource "cloudautomator_post_process" "test" {
					name = "%s"
					group_id = "%s"
					service = "webhook"

					webhook_parameters {
						webhook_authorization_header = "test-authorization-header"
						webhook_url = "http://example.com"
					}
				}`, resourceName, acctest.TestGroupId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "service", "webhook"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "webhook_parameters.0.webhook_authorization_header", "test-authorization-header"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "webhook_parameters.0.webhook_url", "http://example.com"),
			},
		},
		{
			name:            "SharedByGroupTrue",
			postProcessName: fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
				resource "cloudautomator_post_process" "test" {
					name = "%s"
					service = "email"
					shared_by_group = true

					email_parameters {
						email_recipient = "test@example.com"
					}
				}`, resourceName)
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "service", "email"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "shared_by_group", "true"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "email_parameters.0.email_recipient", "test@example.com"),
			},
		},
		{
			name:            "SharedByGroupFalse",
			postProcessName: fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
				resource "cloudautomator_post_process" "test" {
					name = "%s"
					group_id = "%s"
					service = "email"
					shared_by_group = false

					email_parameters {
						email_recipient = "test@example.com"
					}
				}`, resourceName, acctest.TestGroupId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "service", "email"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "shared_by_group", "false"),
				resource.TestCheckResourceAttr("cloudautomator_post_process.test", "email_parameters.0.email_recipient", "test@example.com"),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(tc.postProcessName),
						Check: resource.ComposeTestCheckFunc(
							append([]resource.TestCheckFunc{
								testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], "cloudautomator_post_process.test"),
								resource.TestCheckResourceAttr("cloudautomator_post_process.test", "name", tc.postProcessName),
							}, tc.checks...)...,
						),
					},
				},
			})
		})
	}
}

func testAccCheckCloudAutomatorPostProcessExists(_ *schema.Provider, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		c := testAccProvider.Meta().(*client.Client)

		if err := cloudautomatorPostProcessExistsHelper(s, c, n); err != nil {
			return err
		}

		return nil
	}
}

func cloudautomatorPostProcessExistsHelper(s *terraform.State, c *client.Client, name string) error {
	id := s.RootModule().Resources[name].Primary.ID
	if _, _, err := c.GetPostProcess(id); err != nil {
		return fmt.Errorf("received an error retrieving post process %s", err)
	}

	return nil
}

func testAccCheckCloudAutomatorPostProcessDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudautomator_post_process" {
			continue
		}

		postProcessId := rs.Primary.ID

		_, res, err := c.GetPostProcess(postProcessId)
		if err != nil {
			if res.StatusCode == 404 {
				continue
			}

			return fmt.Errorf("received an error retrieving post process %s", err)
		}

		return fmt.Errorf("post process exists.")
	}

	return nil
}
