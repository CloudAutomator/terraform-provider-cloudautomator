package acctest

import (
	"os"
)

func TestGroupId() string {
	return os.Getenv("CA_TEST_GROUP_ID")
}

func TestAwsAccountId() string {
	return os.Getenv("CA_TEST_AWS_ACCOUNT_ID")
}

func TestSqsAwsAccountId() string {
	return os.Getenv("CA_TEST_SQS_AWS_ACCOUNT_ID")
}

func TestGoogleCloudAccountId() string {
	return os.Getenv("CA_TEST_GOOGLE_CLOUD_ACCOUNT_ID")
}

func TestSqsRegion() string {
	return os.Getenv("CA_TEST_SQS_REGION")
}

func TestSqsQueue() string {
	return os.Getenv("CA_TEST_SQS_QUEUE")
}

func TestPostProcessId() string {
	return os.Getenv("CA_TEST_POST_PROCESS_ID")
}
