pipeline {
    agent any

    environment {
        BINARY_NAME = "myapp"
        DIST_FOLDER = "dist"
        ZIP_NAME = "myapp.zip"
    }

    stages {
        stage('Build Go Binary') {
            steps {
                sh '''
                    mkdir -p ${DIST_FOLDER}
                    go mod tidy
                    go build -o ${DIST_FOLDER}/${BINARY_NAME} main.go
                '''
            }
        }

        stage('Zip Binary') {
            steps {
                sh '''
                    cd ${DIST_FOLDER}
                    zip ${ZIP_NAME} ${BINARY_NAME}
                '''
            }
        }

        stage('Archive ZIP') {
            steps {
                archiveArtifacts artifacts: 'dist/myapp.zip', allowEmptyArchive: false
            }
        }
    }
}
