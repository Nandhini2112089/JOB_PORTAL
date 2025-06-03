FROM jenkins/jenkins:lts

USER root

RUN apt-get update && \
    apt-get install -y curl git unzip && \
    curl -OL https://golang.org/dl/go1.20.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz && \
    rm go1.20.5.linux-amd64.tar.gz


RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.53.3

ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

USER jenkins
