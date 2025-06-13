# jenkins-golang.Dockerfile
FROM jenkins/jenkins:lts

USER root

# Install build essentials and Go 1.22.3
RUN apt-get update && \
    apt-get install -y wget unzip curl zip git && \
    wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz && \
    rm go1.22.3.linux-amd64.tar.gz && \
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile

ENV PATH="/usr/local/go/bin:$PATH"

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
      | sh -s -- -b /usr/local/bin v1.55.2 && \
    golangci-lint --version

USER jenkins
