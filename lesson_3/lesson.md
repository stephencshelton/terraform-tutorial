# Exercise: Using Local Values in Terraform

## Requirements
- Create a new Terraform file named `locals.tf` in the same directory as this lesson.
- Define an AWS provider. Set the region to `us-west-2`.
- Define a local value named `common_tags` that holds a map of tags (key-value pairs). For instance, you can include tags like `Environment`, `Name`, `Project`, etc.
- Create an `aws_instance` resource. Use the `ami` "ami-0c55b159cbfafe1f0" and `t2.micro` for the `instance_type`. Also, assign the tags using the `common_tags` local value you defined.
- After writing the script, run `terraform init` to initialize your Terraform configuration.
- Once initialized, run `terraform apply` to create the EC2 instance with the assigned tags. You should be able to see the instance and its tags in your AWS Console.
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