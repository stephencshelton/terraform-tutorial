package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEC2InstancePublicIp(t *testing.T) {
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

	// Get output value
	outputIp := terraform.Output(t, terraformOptions, "instance_public_ip")

	// Get Instance public IP
	instanceIp := aws.GetPublicIpOfEc2Instance(t, "<instance_id>", awsRegion) // Replace <instance_id> with your actual instance ID

	// Assert public IP
	assert.Equal(t, outputIp, instanceIp)
}