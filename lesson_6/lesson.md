# Exercise: Using for_each to Create Multiple Resources in Terraform

## Requirements
- Create a new Terraform file named `for_each.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Define a local variable named `instance_details` that holds a map of maps. The outer map's key is the instance name and the inner map contains `ami` and `instance_type` details.
- Create an `aws_instance` resource using `for_each` to iterate over the `instance_details` local variable. This should create an instance for each element in the `instance_details`.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the EC2 instances. You should be able to see the instances in your AWS Console.
- Finally, destroy the resources after you're done checking your work using `terraform destroy` to avoid unnecessary AWS charges.

## Requirements for the `instance_details` map
- Include at least two entries in the map.
- Each entry should have a unique name as the key.
- The value of each entry should be a map with `ami` and `instance_type` keys.
- The `ami` should be set to "ami-0c55b159cbfafe1f0".
- The `instance_type` should be set to "t2.micro".

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
