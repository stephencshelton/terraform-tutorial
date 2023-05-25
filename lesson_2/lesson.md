# Exercise: Using Variables in Terraform

## Requirements
- Create a new Terraform file in the same directory as this lesson. Name the file `variables.tf`.
- Define a variable named `instance_type` with a description and default value.
- Define an AWS provider. Set the region to `us-west-2`.
- Create an `aws_instance` resource. You should utilize the `ami` "ami-0c55b159cbfafe1f0" and use the `instance_type` variable you defined earlier for the `instance_type`.
- After writing the script, ensure to run `terraform init` in the terminal to initialize your Terraform configuration.
- After initializing, run `terraform apply` to create the EC2 instance. You should be able to see the instance in your AWS Console.
- Don't forget to destroy the resources after you're done checking your work using `terraform destroy` to avoid unnecessary AWS charges.

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