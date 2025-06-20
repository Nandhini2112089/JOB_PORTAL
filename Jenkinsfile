// // pipeline {
// //     agent any

// //     environment {
// //         BINARY_NAME = "myapp"
//         DIST_FOLDER = "dist"
//         ZIP_NAME = "myapp.zip"
// //     }

// //     stages {
// //         stage('Build Go Binary') {
// //             steps {
// //                 sh '''
// //                     mkdir -p ${DIST_FOLDER}
// //                     go mod tidy
// //                     go build -o ${DIST_FOLDER}/${BINARY_NAME} main.go
// //                 '''
// //             }
// //         }

//         // stage('Zip Binary') {
//         //     steps {
//         //         sh '''
//         //             cd ${DIST_FOLDER}
//         //             zip ${ZIP_NAME} ${BINARY_NAME}
//         //         '''
//         //     }
//         // }

// //         stage('Archive ZIP') {
// //             steps {
// //                 archiveArtifacts artifacts: 'dist/myapp.zip', allowEmptyArchive: false
// //             }
// //         }
// //     }
// // }



pipeline {
    agent any

    environment {
        DIST_FOLDER = "build"
        BINARY_NAME = "app"
        ZIP_NAME    = "job_portal.zip"
        IMAGE_NAME  = "sivanandhini23/db_gorm_app"
        EMAIL_RECIPIENT = "nandhinibalamurugan2003@gmail.com"
        BACKUP_TAG = "previous"
    }

    stages {
        stage('Run Unit Tests') {
            steps {
                sh '''
                export PATH=$PATH:/usr/local/go/bin
                go mod tidy
                go test ./... -v -cover
                '''
            }
        }

        stage('Build Go Binary') {
            steps {
                sh '''
                export PATH=$PATH:/usr/local/go/bin
                mkdir -p build
                go build -o build/app main.go
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
                archiveArtifacts artifacts: 'build/job_portal.zip', allowEmptyArchive: false
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                docker build -t $IMAGE_NAME:latest .
                '''
            }
        }

        stage('Tag Previous Image') {
            steps {
                sh '''
                docker pull $IMAGE_NAME:latest || true
                docker tag $IMAGE_NAME:latest $IMAGE_NAME:$BACKUP_TAG || true
                '''
            }
        }

        stage('Push to Docker Hub') {
            steps {
                sh '''
                docker login -u sivanandhini23 -p Nandhini@23
                docker push $IMAGE_NAME:latest
                docker push $IMAGE_NAME:$BACKUP_TAG || true
                '''
            }
        }
    }

    post {
        success {
            script {
                sleep 10
                emailext(
                    to: "${EMAIL_RECIPIENT}",
                    subject: "${env.JOB_NAME} - Build #${env.BUILD_NUMBER} - SUCCESS",
                    body: """
Build Successful!

Job: ${env.JOB_NAME}
Build Number: ${env.BUILD_NUMBER}

View Logs: ${env.BUILD_URL}
""",
                    mimeType: 'text/plain'
                )
            }
        }

        failure {
            script {
                sleep 10
                sh '''
                echo "Build failed. Rolling back to previous stable image..."
                docker pull $IMAGE_NAME:$BACKUP_TAG || echo "No previous image to roll back"
                docker tag $IMAGE_NAME:$BACKUP_TAG $IMAGE_NAME:latest || true
                '''
                emailext(
                    to: "${EMAIL_RECIPIENT}",
                    subject: "${env.JOB_NAME} - Build #${env.BUILD_NUMBER} - FAILED & ROLLED BACK",
                    body: """
Build Failed and rolled back to previous stable Docker image.

Job: ${env.JOB_NAME}
Build Number: ${env.BUILD_NUMBER}

Console Output: ${env.BUILD_URL}
""",
                    mimeType: 'text/plain'
                )
            }
        }
    }
}


