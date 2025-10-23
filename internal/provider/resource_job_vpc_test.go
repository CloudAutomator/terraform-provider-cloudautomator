package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Vpc(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "CreateNatGatewayAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_nat_gateway"
				create_nat_gateway_action_value {
				  region           = "ap-northeast-1"
				  allocation_id    = "eipalloc-0123456789abcdef0"
				  nat_gateway_name = "test-nat-gateway"
				  subnet_id        = "subnet-0123456789abcdef0"
				  route_table_id   = "rtb-0123456789abcdef0"

				  additional_tags {
					key   = "key1"
					value = "value1"
				  }

				  additional_tags {
					key   = "key2"
					value = "value2"
				  }

				  additional_tags {
					key   = "key3"
					value = "value3"
				  }
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_nat_gateway"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.allocation_id", "eipalloc-0123456789abcdef0"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.nat_gateway_name", "test-nat-gateway"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.subnet_id", "subnet-0123456789abcdef0"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.route_table_id", "rtb-0123456789abcdef0"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.0.key", "key1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.0.value", "value1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.1.key", "key2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.1.value", "value2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.2.key", "key3"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_nat_gateway_action_value.0.additional_tags.2.value", "value3"),
			},
		},
		{
			name:    "DeleteNatGatewayAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "delete_nat_gateway"
				delete_nat_gateway_action_value {
				  region    = "ap-northeast-1"
				  tag_key   = "Name"
				  tag_value = "test-nat-gateway"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "delete_nat_gateway"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_nat_gateway_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_nat_gateway_action_value.0.tag_key", "Name"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "delete_nat_gateway_action_value.0.tag_value", "test-nat-gateway"),
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			baseChecks := []resource.TestCheckFunc{
				testAccCheckCloudAutomatorJobExists(testAccProviders["cloudautomator"], "cloudautomator_job.test"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "name", tc.jobName),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "completed_post_process_id.0", acctest.TestPostProcessId()),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "failed_post_process_id.0", acctest.TestPostProcessId()),
			}

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testAccCheckCloudAutomatorJobDestroy,
				Steps: []resource.TestStep{
					{
						Config: tc.configFunc(tc.jobName),
						Check: resource.ComposeAggregateTestCheckFunc(
							append(baseChecks, tc.checks...)...,
						),
					},
				},
			})
		})
	}
}
