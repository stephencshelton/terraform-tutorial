package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVPCCreationWithCommunityModule(t *testing.T) {
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

	// Get the VPC ID
	vpcID := terraform.Output(t, terraformOptions, "vpc_id")

	// Get the Subnets IDs
	subnetsIDs := terraform.OutputList(t, terraformOptions, "subnet_ids")

	// Check if the VPC and the Subnets exist
	assert.True(t, aws.IsVpcExist(t, awsRegion, vpcID))
	assert.Equal(t, 2, len(subnetsIDs))
	for _, subnetID := range subnetsIDs {
		assert.True(t, aws.IsSubnetExist(t, awsRegion, subnetID))
	}
}