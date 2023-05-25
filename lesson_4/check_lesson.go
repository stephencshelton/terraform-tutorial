package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEC2InstanceCreationWithDataSource(t *testing.T) {
	// AWS Region
	awsRegion := "us-west-2"

	// Terraform options
	terraformOptions := &terraform.Options{
		// Terraform directory
		TerraformDir: "<terraform_directory>", // Replace <terraform_directory> with the actual directory path
	}

	// Destroy resources after test
	defer terraform.Destroy(t, terraformOptions)

	// Init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Define expected instance
	expectedInstance := "<instance_id>" // Replace with your actual instance ID

	// Check if instance exists
	exists := aws.DoesEC2InstanceExist(t, expectedInstance, awsRegion)

	// Assert instance existence
	assert.True(t, exists)
}