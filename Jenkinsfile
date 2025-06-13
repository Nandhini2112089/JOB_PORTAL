pipeline {
    agent any

    environment {
        GOPROXY = 'https://proxy.golang.org,direct'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Download Dependencies') {
            steps {
                sh 'go mod download'
            }
        }

        stage('Clean Cache') {
            steps {
                sh 'go clean -modcache'
            }
        }

        // Removed Lint stage

        stage('Unit Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build Binary') {
            steps {
                sh 'go build -o myapp'
            }
        }

        stage('Zip Binary') {
            steps {
                sh 'zip myapp.zip myapp'
            }
        }

        stage('Archive Artifact') {
            steps {
                archiveArtifacts artifacts: 'myapp.zip', fingerprint: true
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished.'
        }
        failure {
            echo 'Pipeline failed.'
        }
    }
}
