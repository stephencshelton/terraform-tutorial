# Exercise: Using Conditional Expressions in Terraform

## Requirements
- Create a new Terraform file named `conditionals.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Define a variable named `create_instance` with a default value of `true`.
- Create an `aws_instance` resource using the `ami` "ami-0c55b159cbfafe1f0" and `t2.micro` for the `instance_type`. However, the creation of this resource should depend on the value of `create_instance`. If `create_instance` is `false`, no instance should be created.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to potentially create the EC2 instance, depending on your `create_instance` variable. You should be able to see the instance in your AWS Console if `create_instance` is `true`.
- Finally, if an instance was created, destroy it after you're done checking your work using `terraform destroy` to avoid unnecessary AWS charges.

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