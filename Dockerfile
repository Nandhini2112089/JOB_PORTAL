FROM jenkins/jenkins:lts

USER root

# Install Go 1.21 manually
RUN apt-get update && \
    apt-get install -y wget zip && \
    wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz && \
    rm go1.21.6.linux-amd64.tar.gz

# Set PATH for Go
ENV PATH="/usr/local/go/bin:${PATH}"

USER jenkins
