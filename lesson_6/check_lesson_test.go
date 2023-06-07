package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEC2InstancesCreationWithForEach(t *testing.T) {
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

	// Define expected instances
	expectedInstances := []string{"<instance_id_1>", "<instance_id_2>"} // Replace with your actual instance IDs

	// Loop over the expected instances and assert their existence
	for _, instanceId := range expectedInstances {
		// Check if instance exists
		exists := aws.DoesEC2InstanceExist(t, instanceId, awsRegion)

		// Assert instance existence
		assert.True(t, exists)
	}
}
