package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestConditionalEC2Creation(t *testing.T) {
	// AWS Region
	awsRegion := "us-west-2"

	// Terraform options
	terraformOptions := &terraform.Options{
		// Terraform directory
		TerraformDir: "<terraform_directory>", // Replace <terraform_directory> with the actual directory path

		// Variables
		Vars: map[string]interface{}{
			"create_instance": true,
		},
	}

	// Destroy resources after test
	defer terraform.Destroy(t, terraformOptions)

	// Init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Check if instance exists
	exists := aws.DoesEC2InstanceExist(t, "<instance_id>", awsRegion) // Replace <instance_id> with your actual instance ID

	// Assert instance existence
	assert.True(t, exists)
}