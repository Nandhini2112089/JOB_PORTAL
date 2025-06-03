pipeline {
    agent any

    environment {
        GO111MODULE = 'on'
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.GOPATH}/bin:/usr/local/go/bin:/usr/local/bin:${env.PATH}"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Prepare') {
            steps {
                // Download dependencies and build to generate export data needed by golangci-lint
                sh 'go clean -modcache || true'
                sh 'go mod download'
                sh 'go build ./...'
            }
        }

        stage('Lint') {
            steps {
                sh 'golangci-lint run ./... --timeout 5m'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -v'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myapp cmd/main.go'
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'myapp'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
