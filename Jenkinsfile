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
        DIST_FOLDER = "binary"
        ZIP_NAME = "job_portal.zip"
        IMAGE_NAME = 'sivanandhini23/db_gorm_app'
        EMAIL_RECIPIENT = 'nandhinibalamurugan2003@gmail.com'
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
                archiveArtifacts artifacts: 'binary/job_portal.zip', allowEmptyArchive: false
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME:latest .'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                 sh '''
                docker login -u sivanandhini23 -p Nandhini@23
                docker push $IMAGE_NAME:latest
                '''
               
            }
        }
    }

    post {
        success {
            script {
                sleep 10 // delay to ensure SMTP readiness
                emailext(
                    to: "${EMAIL_RECIPIENT}",
                    subject: "${env.JOB_NAME} - Build #${env.BUILD_NUMBER} - SUCCESS",
                    body: """
 Build Successful!

 Job: ${env.JOB_NAME}  
 Build Number: ${env.BUILD_NUMBER}  

 View Build Output: ${env.BUILD_URL}
""",
                    mimeType: 'text/plain'
                )
            }
        }

        failure {
            script {
                sleep 10
                emailext(
                    to: "${EMAIL_RECIPIENT}",
                    subject: "${env.JOB_NAME} - Build #${env.BUILD_NUMBER} - FAILED",
                    body: """
 Build Failed!

 Job: ${env.JOB_NAME}  
 Build Number: ${env.BUILD_NUMBER}  

 Check console output: ${env.BUILD_URL}
""",
                    mimeType: 'text/plain'
                )
            }
        }
    }
}

