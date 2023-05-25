package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEC2InstanceTags(t *testing.T) {
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

	// Define expected tags
	expectedTags := map[string]string{
		"Environment": "Dev",
		"Name":        "MyInstance",
		"Project":     "LearningTerraform",
	}

	// Get Instance tags
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, "<instance_id>") // Replace <instance_id> with your actual instance ID

	// Assert tags
	assert.Equal(t, expectedTags, instanceTags)
}