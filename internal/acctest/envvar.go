package acctest

import (
	"os"
)

func TestGroupId() string {
	return os.Getenv("CA_TEST_GROUP_ID")
}

func TestAwsAccountNumber() string {
	return os.Getenv("CA_TEST_AWS_ACCOUNT_NUMBER")
}

func TestAwsAccountId() string {
	return os.Getenv("CA_TEST_AWS_ACCOUNT_ID")
}

func TestGoogleCloudAccountId() string {
	return os.Getenv("CA_TEST_GOOGLE_CLOUD_ACCOUNT_ID")
}

func TestRegion() string {
	return os.Getenv("CA_TEST_REGION")
}

func TestS3BucketName() string {
	return os.Getenv("CA_TEST_S3_BUCKET_NAME")
}

func TestSqsQueue() string {
	return os.Getenv("CA_TEST_SQS_QUEUE")
}

func TestPostProcessId() string {
	return os.Getenv("CA_TEST_POST_PROCESS_ID")
}
