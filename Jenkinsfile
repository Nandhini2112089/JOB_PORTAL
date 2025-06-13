pipeline {
    agent any

    environment {
        GO_BIN = "/usr/local/go/bin"
        PATH = "${GO_BIN}:${env.PATH}"
        DEST = "${WORKSPACE}/artifact_output"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Lint') {
            steps {
                sh 'golangci-lint run ./...'
            }
        }

        stage('Unit Test') {
            steps {
                sh 'go test ./... -v -cover'
            }
        }

        stage('Build Binary') {
            steps {
                sh 'go build -o main main.go'
            }
        }

        stage('Zip Binary') {
            steps {
                sh '''
                    mkdir -p ${DEST}
                    cp main ${DEST}/
                    cd ${DEST}
                    zip -r go_grpc_app.zip main
                '''
            }
        }

        stage('Archive Artifact') {
            steps {
                archiveArtifacts artifacts: 'artifact_output/go_grpc_app.zip', allowEmptyArchive: false
            }
        }
    }

    post {
        success {
            echo 'Pipeline completed successfully.'
        }
        failure {
            echo 'Pipeline failed.'
        }
    }
}
