package cloudautomator

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const testApiKeyEnvName = "CA_TEST_API_KEY"
const testAwsAccountIdEnvName = "CA_TEST_AWS_ACCOUNT_ID"
const testGroupIdEnvName = "CA_TEST_GROUP_ID"
const testSqsAwsAccountIdEnvName = "CA_TEST_SQS_AWS_ACCOUNT_ID"
const testSqsRegionEnvName = "CA_TEST_SQS_REGION"
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

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
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

	if err := os.Getenv(testSqsAwsAccountIdEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testSqsAwsAccountIdEnvName)
	}

	if err := os.Getenv(testSqsRegionEnvName); err == "" {
		t.Fatalf("%s must be set for acceptance tests", testSqsRegionEnvName)
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

	if err := os.Setenv(ApiKeyEnvName, os.Getenv(testApiKeyEnvName)); err != nil {
		t.Fatalf("Error setting API key: %v", err)
	}
}
