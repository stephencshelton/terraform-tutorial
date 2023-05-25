# Exercise: Using Provisioners in Terraform

## Requirements

- Create a new Terraform file in the same directory as this script.
- Name the file `main.tf`.
- Create a null_resource.
- Add a local-exec provisioner to the null_resource that runs a simple command like `echo "Hello, World!"`.
- Apply your configuration and make sure the provisioner runs as expected.

**NOTE**: Remember to run `terraform init` before applying your configuration.

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