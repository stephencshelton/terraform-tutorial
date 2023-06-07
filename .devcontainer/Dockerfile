# Use the official image as a parent image
FROM ubuntu:22.04

# Set environment variables for versions
ENV GO_VERSION 1.20.5
ENV TERRAFORM_VERSION 1.4.6
ENV TERM xterm-256color

# Set the working directory
WORKDIR /workspace

# Install necessary packages
RUN apt-get update && apt-get install -y \
    git \
    curl \
    unzip \
    build-essential \
    dos2unix \
    python3-pip \
    bash-completion

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

# Add tab completion, colors, and a fancy prompt for the vscode user
USER vscode
RUN echo 'source /etc/profile' >> /home/vscode/.bashrc \
    && echo 'complete -C '/usr/local/bin/aws_completer' aws' >> /home/vscode/.bashrc \
    && echo "export PS1=\"\[\e[36m\]\u\[\e[m\]@\[\e[32m\]\h:\[\e[33;1m\]\W\[\e[m\]\$ \"" >> /home/vscode/.bashrc \
    && echo "alias ls='ls --color=auto'" >> /home/vscode/.bashrc \
    && echo "alias ll='ls -l'" >> /home/vscode/.bashrc \
    && echo "alias l='ls -CF'" >> /home/vscode/.bashrc \
    && echo "alias la='ls -A'" >> /home/vscode/.bashrc \
    && echo "alias lla='ls -lA'" >> /home/vscode/.bashrc \
    && echo "export PATH=/usr/local/go/bin:\$PATH" >> /home/vscode/.bashrc && \
    && terraform -install-autocomplete

# Copy the content of your repository into the container
USER root
COPY --chown=vscode:vscode . .
RUN dos2unix ./*
WORKDIR /home/vscode

# This will enable the terminal's interactive shell
CMD [ "/bin/bash" ]