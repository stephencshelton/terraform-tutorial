# Exercise: Using Community Modules in Terraform

## Requirements
- Create a new Terraform file named `community_module.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Use the Terraform AWS VPC Module to create a VPC. You can find this module at [Terraform AWS VPC Module](https://registry.terraform.io/modules/terraform-aws-modules/vpc/aws/latest).
- Customize the inputs of the module to create a VPC with 2 public subnets.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the VPC and the subnets.
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