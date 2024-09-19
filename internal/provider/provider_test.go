package cloudautomator

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testApiKeyEnvName = "CA_TEST_API_KEY"
const testAwsAccountNumberEnvName = "CA_TEST_AWS_ACCOUNT_NUMBER"
const testAwsAccountIdEnvName = "CA_TEST_AWS_ACCOUNT_ID"
const testGoogleCloudAccountIdEnvName = "CA_TEST_GOOGLE_CLOUD_ACCOUNT_ID"
const testGroupIdEnvName = "CA_TEST_GROUP_ID"
const testRegionEnvName = "CA_TEST_REGION"
const TestS3BucketNameEnvName = "CA_TEST_S3_BUCKET_NAME"
const testSqsQueueEnvName = "CA_TEST_SQS_QUEUE"
const testPostProcessIdEnvName = "CA_TEST_POST_PROCESS_ID"

var testAccProviderFactories map[string]func() (*schema.Provider, error)
var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"cloudautomator": testAccProvider,
	}
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"cloudautomator": func() (*schema.Provider, error) { return testAccProvider, nil },
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv(testApiKeyEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testApiKeyEnvName)
	}

	if err := os.Getenv(testGroupIdEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testGroupIdEnvName)
	}

	if err := os.Getenv(testAwsAccountIdEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testAwsAccountIdEnvName)
	}

	if err := os.Getenv(testAwsAccountNumberEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testAwsAccountNumberEnvName)
	}

	if err := os.Getenv(testGoogleCloudAccountIdEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testGoogleCloudAccountIdEnvName)
	}

	if err := os.Getenv(testRegionEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testRegionEnvName)
	}

	if err := os.Getenv(testSqsQueueEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testSqsQueueEnvName)
	}

	if err := os.Getenv(testPostProcessIdEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testPostProcessIdEnvName)
	}

	if err := os.Getenv(testApiKeyEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testApiKeyEnvName)
	}

	if err := os.Setenv(apiKeyEnvName, os.Getenv(testApiKeyEnvName)); err != nil {
		t.Fatalf("Error setting API key: %v", err)
	}
}
