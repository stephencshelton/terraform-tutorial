# Exercise: Using Outputs in Terraform

## Requirements
- Create a new Terraform file named `outputs.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Create an `aws_instance` resource. You should utilize the `ami` "ami-0c55b159cbfafe1f0" and use `t2.micro` for the `instance_type`.
- Define an output named `instance_public_ip` that retrieves the `public_ip` of the `aws_instance` created.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the EC2 instance and display the instance's public IP. You should be able to see the instance and its public IP in your AWS Console.
- Finally, destroy the resources after you're done checking your work using `terraform destroy` to avoid unnecessary AWS charges.

## Hint
- Remember, you can get the `public_ip` of an `aws_instance` like this: `aws_instance.<resource_name>.public_ip`.

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