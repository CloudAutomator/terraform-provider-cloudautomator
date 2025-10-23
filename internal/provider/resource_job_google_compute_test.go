package cloudautomator

import (
	"fmt"
	"testing"

	"terraform-provider-cloudautomator/internal/acctest"
	"terraform-provider-cloudautomator/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudAutomatorJob_GoogleCompute(t *testing.T) {
	cases := []struct {
		name       string
		jobName    string
		configFunc func(string) string
		checks     []resource.TestCheckFunc
	}{
		{
			name:    "GoogleComputeInsertMachineImageAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
				name = "%s"
				group_id = "%s"
				google_cloud_account_id = "%s"

				rule_type = "webhook"

				action_type = "google_compute_insert_machine_image"
				google_compute_insert_machine_image_action_value {
					region = "asia-northeast1"
					project_id = "example-project"
					specify_vm_instance = "label"
					vm_instance_label_key = "env"
					vm_instance_label_value = "develop"
					machine_image_storage_location = "asia-northeast1"
					machine_image_basename = "example-daily"
					generation = 10
				}
				completed_post_process_id = [%s]
				failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "google_compute_insert_machine_image"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.region", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.project_id", "example-project"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.specify_vm_instance", "label"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.vm_instance_label_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.vm_instance_label_value", "develop"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.machine_image_storage_location", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.machine_image_basename", "example-daily"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_insert_machine_image_action_value.0.generation", "10"),
			},
		},
		{
			name:    "GoogleComputeStartVmInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
			  name = "%s"
			  group_id = "%s"
			  google_cloud_account_id = "%s"

			  rule_type = "webhook"

			  action_type = "google_compute_start_vm_instances"
			  google_compute_start_vm_instances_action_value {
				region = "asia-northeast1"
				project_id = "example-project"
				specify_vm_instance = "label"
				vm_instance_label_key = "env"
				vm_instance_label_value = "develop"
			  }
			  completed_post_process_id = [%s]
			  failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "google_compute_start_vm_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.region", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.project_id", "example-project"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.specify_vm_instance", "label"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.vm_instance_label_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_start_vm_instances_action_value.0.vm_instance_label_value", "develop"),
			},
		},
		{
			name:    "GoogleComputeStopVmInstancesAction",
			jobName: fmt.Sprintf("tf-testacc-job-%s", utils.RandomString(12)),
			configFunc: func(resourceName string) string {
				return fmt.Sprintf(`
			resource "cloudautomator_job" "test" {
			  name = "%s"
			  group_id = "%s"
			  google_cloud_account_id = "%s"

			  rule_type = "webhook"

			  action_type = "google_compute_stop_vm_instances"
			  google_compute_stop_vm_instances_action_value {
				region = "asia-northeast1"
				project_id = "example-project"
				specify_vm_instance = "label"
					vm_instance_label_key = "env"
				vm_instance_label_value = "develop"
			  }
			  completed_post_process_id = [%s]
			  failed_post_process_id = [%s]
			}`, resourceName, acctest.TestGroupId(), acctest.TestGoogleCloudAccountId(), acctest.TestPostProcessId(), acctest.TestPostProcessId())
			},
			checks: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("cloudautomator_job.test", "action_type", "google_compute_stop_vm_instances"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.region", "asia-northeast1"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.project_id", "example-project"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.specify_vm_instance", "label"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.vm_instance_label_key", "env"),
				resource.TestCheckResourceAttr("cloudautomator_job.test", "google_compute_stop_vm_instances_action_value.0.vm_instance_label_value", "develop"),
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
