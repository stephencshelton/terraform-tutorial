# Exercise: Using Data Sources in Terraform

## Requirements
- Create a new Terraform file named `data_sources.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Use a data source to retrieve the latest Ubuntu Server 20.04 LTS AMI that is EBS-backed, uses the HVM virtualization type, and is owned by Amazon.
- Create an `aws_instance` resource using the `ami` retrieved from the data source and `t2.micro` for the `instance_type`.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the EC2 instance. You should be able to see the instance in your AWS Console.
- Finally, destroy the resources after you're done checking your work using `terraform destroy` to avoid unnecessary AWS charges.

## Checking Your Work

Once you have completed the exercise, you can check your work by running the corresponding Terratest in the devcontainer.

1. Open a terminal in VSCode (Terminal > New Terminal).
2. Navigate to the directory containing the lesson files.
3. Run the following commands:

```bash
# Download dependencies
go mod init check
go get github.com/gruntwork-io/terratest/modules/terraform
go get github.com/stretchr/testify/assert

# Run the test
go test -v