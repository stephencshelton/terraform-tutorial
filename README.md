# Terraform Learning Repository

This repository provides an interactive learning experience for Terraform. Each lesson has a Terraform configuration file (.tf) with commented instructions to follow and a corresponding Terratest to validate the implementation.

## Requirements

- Visual Studio Code (VSCode)
- Docker Desktop
- Windows Subsystem for Linux 2 (WSL2)
- Devcontainer extension for VSCode

## Setup

1. **Install Visual Studio Code**: Download and install VSCode from the [official website](https://code.visualstudio.com/).

2. **Install the devcontainer extension for VSCode**: Follow the directions from the [official website](https://code.visualstudio.com/docs/remote/containers).

3. **Install Docker Desktop**: Download and install Docker Desktop from the [official website](https://www.docker.com/products/docker-desktop).

4. **Setup WSL2 for Docker Desktop**:
    - Install WSL2 by following the instructions from the [Microsoft's official guide](https://docs.microsoft.com/en-us/windows/wsl/install-win10).
    - After installing WSL2, go to Docker Desktop Settings -> Resources -> WSL Integration, and enable integration with your Linux distro.

5. **Setup AWS account and access keys**:
    - Create an AWS account if you don't already have one.
    - Create access keys from your AWS account. To create new access keys, go to your AWS Management Console > Security Credentials > Access Keys (Access Key ID and Secret Access Key) > Create New Access Key.
    - Make note of the Access Key ID and Secret Access Key. You will need these to set up the AWS CLI.

6. **Clone this repository**: Clone this repository to your local machine and open the cloned folder in VSCode.

7. **Open repository in devcontainer**: Once you clone the repository hit `F1` and type `dev` which will show the option `Dev Containers: Rebuild and Reopen in Container` which will launch your ubuntu environment to write/run your Terraform configuration and Terratests.

8. **Configure AWS CLI**: Inside the devcontainer terminal, type `aws configure` and follow the prompts. Enter the Access Key ID and Secret Access Key you noted down earlier. You can set the default region to `us-west-2` (Oregon) and the default output format to `json`.

## Lessons Overview

**Terraform Lessons:**

- **Lesson 1**: Learn to use provisioners in Terraform.
- **Lesson 2**: Learn to use variables in Terraform.
- **Lesson 3**: Learn to use locals in Terraform.
- **Lesson 4**: Learn to use data sources in Terraform.
- **Lesson 5**: Learn to use outputs in Terraform.
- **Lesson 6**: Learn to use `for_each` loop in Terraform.
- **Lesson 7**: Learn to use `count` parameter in Terraform.
- **Lesson 8**: Learn to use Terraform modules.
- **Lesson 9**: Learn to use conditional expressions in Terraform.
- **Lesson 10**: Learn to use community modules in Terraform.