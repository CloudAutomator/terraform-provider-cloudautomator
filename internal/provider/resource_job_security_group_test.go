package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_SecurityGroup(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "AuthorizeSecurityGroupIngressAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "authorize_security_group_ingress"
				authorize_security_group_ingress_action_value {
					region = "ap-northeast-1"
					specify_security_group = "tag"
					tag_key = "env"
					tag_value = "develop"
					ip_protocol = "tcp"
					to_port = "80"
					cidr_ip = "172.31.0.0/16"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "authorize_security_group_ingress"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.specify_security_group", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.ip_protocol", "tcp"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.to_port", "80"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "authorize_security_group_ingress_action_value.0.cidr_ip", "172.31.0.0/16"),
			},
		},
		{
			name:    "RevokeSecurityGroupIngressAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "revoke_security_group_ingress"
				revoke_security_group_ingress_action_value {
					region = "ap-northeast-1"
					specify_security_group = "tag"
					tag_key = "env"
					tag_value = "develop"
					ip_protocol = "tcp"
					to_port = "80"
					cidr_ip = "172.31.0.0/16"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "revoke_security_group_ingress"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.specify_security_group", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.ip_protocol", "tcp"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.to_port", "80"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "revoke_security_group_ingress_action_value.0.cidr_ip", "172.31.0.0/16"),
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
