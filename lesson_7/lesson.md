# Exercise: Using `count` in Terraform

## Requirements
- Create a new Terraform file named `count.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Create an `aws_instance` resource. Set the `ami` to "ami-0c55b159cbfafe1f0" and the `instance_type` to "t2.micro". 
- Use the `count` parameter to create 3 instances of the same type.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the EC2 instances. You should be able to see 3 identical instances in your AWS Console.
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