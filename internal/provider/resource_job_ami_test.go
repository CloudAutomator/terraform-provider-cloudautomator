package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_Ami(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "CreateImageAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "create_image"
				create_image_action_value {
					region = "ap-northeast-1"
					specify_image_instance = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					image_name = "test-image"
					description = "test image"
					reboot_instance = "true"
					additional_tags {
						key = "key-1"
						value = "value-1"
					}
					additional_tags {
						key = "key-2"
						value = "value-2"
					}
					add_same_tag_to_snapshot = "true"
					trace_status = "true"
					recreate_image_if_ami_status_failed	 = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "create_image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.specify_image_instance", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.image_name", "test-image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.description", "test image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.reboot_instance", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.0.key", "key-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.0.value", "value-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.1.key", "key-2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.additional_tags.1.value", "value-2"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.add_same_tag_to_snapshot", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.trace_status", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "create_image_action_value.0.recreate_image_if_ami_status_failed", "true"),
			},
		},
		{
			name:    "CopyImageAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_id = "%s"

				rule_type = "webhook"

				action_type = "copy_image"
				copy_image_action_value {
					source_region = "ap-northeast-1"
					destination_region = "us-east-1"
					specify_image = "tag"
					tag_key = "env"
					tag_value = "develop"
					generation = 10
					trace_status = "true"
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "copy_image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.source_region", "ap-northeast-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.destination_region", "us-east-1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.specify_image", "tag"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.tag_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.tag_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.generation", "10"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "copy_image_action_value.0.trace_status", "true"),
			},
		},
		{
			name:    "BulkDeleteImagesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				aws_account_ids = [%s]

				rule_type = "webhook"

				action_type = "bulk_delete_images"
				bulk_delete_images_action_value {
					exclude_by_tag_bulk_delete_images = true
					exclude_by_tag_key_bulk_delete_images = "env"
					exclude_by_tag_value_bulk_delete_images = "production"
					specify_base_date = "before_days"
					before_days = 365
				}

				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestAwsAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "bulk_delete_images"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.exclude_by_tag_bulk_delete_images", "true"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.exclude_by_tag_key_bulk_delete_images", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.exclude_by_tag_value_bulk_delete_images", "production"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.specify_base_date", "before_days"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "bulk_delete_images_action_value.0.before_days", "365"),
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
