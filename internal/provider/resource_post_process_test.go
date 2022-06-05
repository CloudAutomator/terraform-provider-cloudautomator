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

func TestAccCloudAutomatorPostProcess_Email(t *testing.T) {
	resourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorPostProcessConfigEmail(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						resourceName, "service", "email"),
					resource.TestCheckResourceAttr(
						resourceName, "group_id", acctest.TestGroupId()),
					resource.TestCheckResourceAttr(
						resourceName, "shared_by_group", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "email_parameters.0.email_recipient", "test@example.com"),
				),
			},
		},
	})
}

func TestAccCloudAutomatorPostProcess_Slack(t *testing.T) {
	resourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorPostProcessConfigSlack(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						resourceName, "service", "slack"),
					resource.TestCheckResourceAttr(
						resourceName, "group_id", acctest.TestGroupId()),
					resource.TestCheckResourceAttr(
						resourceName, "shared_by_group", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "slack_parameters.0.slack_channel_name", "ca-test-info-notification"),
					resource.TestCheckResourceAttr(
						resourceName, "slack_parameters.0.slack_language", "ja"),
					resource.TestCheckResourceAttr(
						resourceName, "slack_parameters.0.slack_time_zone", "Tokyo"),
				),
			},
		},
	})
}

func TestAccCloudAutomatorPostProcess_Sqs(t *testing.T) {
	resourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorPostProcessConfigSqs(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						resourceName, "service", "sqs"),
					resource.TestCheckResourceAttr(
						resourceName, "shared_by_group", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "sqs_parameters.0.sqs_aws_account_id", acctest.TestSqsAwsAccountId()),
					resource.TestCheckResourceAttr(
						resourceName, "sqs_parameters.0.sqs_queue", "test-queue"),
					resource.TestCheckResourceAttr(
						resourceName, "sqs_parameters.0.sqs_region", "ap-northeast-1"),
				),
			},
		},
	})
}

func TestAccCloudAutomatorPostProcess_Webhook(t *testing.T) {
	resourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorPostProcessConfigWebhook(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						resourceName, "service", "webhook"),
					resource.TestCheckResourceAttr(
						resourceName, "shared_by_group", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "webhook_parameters.0.webhook_authorization_header", "test-authorization-header"),
					resource.TestCheckResourceAttr(
						resourceName, "webhook_parameters.0.webhook_url", "http://example.com"),
				),
			},
		},
	})
}

func TestAccCloudAutomatorPostProcess_SharedByGroupTrue(t *testing.T) {
	resourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorPostProcessConfigSharedByGroupTrue(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						resourceName, "service", "email"),
					resource.TestCheckResourceAttr(
						resourceName, "shared_by_group", "true"),
					resource.TestCheckResourceAttr(
						resourceName, "email_parameters.0.email_recipient", "test@example.com"),
				),
			},
		},
	})
}

func TestAccCloudAutomatorPostProcess_SharedByGroupFalse(t *testing.T) {
	resourceName := "cloudautomator_post_process.test"
	postProcessName := fmt.Sprintf("tf-testacc-post-process-%s", utils.RandomString(12))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCloudAutomatorPostProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAutomatorPostProcessConfigSharedByGroupFalse(postProcessName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudAutomatorPostProcessExists(testAccProviders["cloudautomator"], resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "name", postProcessName),
					resource.TestCheckResourceAttr(
						resourceName, "service", "email"),
					resource.TestCheckResourceAttr(
						resourceName, "group_id", acctest.TestGroupId()),
					resource.TestCheckResourceAttr(
						resourceName, "shared_by_group", "false"),
					resource.TestCheckResourceAttr(
						resourceName, "email_parameters.0.email_recipient", "test@example.com"),
				),
			},
		},
	})
}

func testAccCheckCloudAutomatorPostProcessExists(accProvider *schema.Provider, n string) resource.TestCheckFunc {
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

func testAccCheckCloudAutomatorPostProcessConfigEmail(rName string) string {
	return fmt.Sprintf(`
resource "cloudautomator_post_process" "test" {
	name = "%s"
	group_id = "%s"
	service = "email"

	email_parameters {
		email_recipient = "test@example.com"
	}
}`, rName, acctest.TestGroupId())
}

func testAccCheckCloudAutomatorPostProcessConfigSlack(rName string) string {
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
}`, rName, acctest.TestGroupId())
}

func testAccCheckCloudAutomatorPostProcessConfigSqs(rName string) string {
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
}`, rName, acctest.TestGroupId(), acctest.TestAwsAccountId())
}

func testAccCheckCloudAutomatorPostProcessConfigWebhook(rName string) string {
	return fmt.Sprintf(`
resource "cloudautomator_post_process" "test" {
	name = "%s"
	group_id = "%s"
	service = "webhook"

	webhook_parameters {
		webhook_authorization_header = "test-authorization-header"
		webhook_url = "http://example.com"
	}
}`, rName, acctest.TestGroupId())
}

func testAccCheckCloudAutomatorPostProcessConfigSharedByGroupTrue(rName string) string {
	return fmt.Sprintf(`
resource "cloudautomator_post_process" "test" {
	name = "%s"
	service = "email"
	shared_by_group = true

	email_parameters {
		email_recipient = "test@example.com"
	}
}`, rName)
}

func testAccCheckCloudAutomatorPostProcessConfigSharedByGroupFalse(rName string) string {
	return fmt.Sprintf(`
resource "cloudautomator_post_process" "test" {
	name = "%s"
	group_id = "%s"
	service = "email"

	email_parameters {
		email_recipient = "test@example.com"
	}
}`, rName, acctest.TestGroupId())
}
