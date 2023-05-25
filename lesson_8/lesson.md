# Exercise: Creating and Using Modules in Terraform

## Requirements
- Create a new directory named `modules` in the same directory as this lesson. Inside the `modules` directory, create another directory named `ec2_instance`.
- In the `ec2_instance` directory, create a new Terraform file named `main.tf`. This file should define a module for creating an EC2 instance. The module should take parameters for `ami` and `instance_type`, and use these to create an `aws_instance` resource.
- In the same directory as this lesson, create a new Terraform file named `modules.tf`. In this file, call the EC2 instance module you defined, providing it with the required parameters. Create at least two EC2 instances using the module.
- Define an AWS provider in the `modules.tf` file. Set the region to `us-west-2`.
- After writing the scripts, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the EC2 instances. You should be able to see the instances in your AWS Console.
- Finally, destroy the resources after you're done checking your work using `terraform destroy` to avoid unnecessary AWS charges.

## Requirements for the EC2 instance module
- The module should define variables for `ami` and `instance_type`.
- The module should create an `aws_instance` resource using the provided `ami` and `instance_type`.
- The `aws_instance` resource should have a unique name for each instance created.

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
