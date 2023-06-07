# Use the official image as a parent image
FROM ubuntu:22.04

# Set the working directory
WORKDIR /workspace

# Set environment variables for versions
ENV GO_VERSION 1.20.5
ENV TERRAFORM_VERSION 1.4.6

# Install necessary packages
RUN apt-get update && apt-get install -y \
    git \
    curl \
    unzip \
    build-essential \
    dos2unix \
    python3-pip

# Install AWS CLI
RUN pip3 install --upgrade awscli

# Install Go
RUN curl -LO https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

# Install AWS CLI
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm awscliv2.zip

# Install Terraform
RUN curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Clean up to reduce image size
RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# This will automatically enable the vscode server's features
RUN useradd -m vscode
USER vscode

# Copy the content of your repository into the container
USER root
COPY --chown=vscode:vscode . .
RUN dos2unix ./*
WORKDIR /home/vscode

# This will enable the terminal's interactive shell
CMD [ "/bin/bash" ]