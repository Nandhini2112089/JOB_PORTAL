// pipeline {
//     agent any

//     environment {
//         BINARY_NAME = "myapp"
//         DIST_FOLDER = "dist"
//         ZIP_NAME = "myapp.zip"
//     }

//     stages {
//         stage('Build Go Binary') {
//             steps {
//                 sh '''
//                     mkdir -p ${DIST_FOLDER}
//                     go mod tidy
//                     go build -o ${DIST_FOLDER}/${BINARY_NAME} main.go
//                 '''
//             }
//         }

//         stage('Zip Binary') {
//             steps {
//                 sh '''
//                     cd ${DIST_FOLDER}
//                     zip ${ZIP_NAME} ${BINARY_NAME}
//                 '''
//             }
//         }

//         stage('Archive ZIP') {
//             steps {
//                 archiveArtifacts artifacts: 'dist/myapp.zip', allowEmptyArchive: false
//             }
//         }
//     }
// }
pipeline {
    agent any

    environment {
        IMAGE_NAME = 'sivanandhini23/db_gorm_app'
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
        emailext(
            to: 'nandhinibalamurugan2003@gmail.com',
            subject: " $JOB_NAME - Build #$BUILD_NUMBER - SUCCESS",
            body: """
Build Successful!

Job: $JOB_NAME  
Build Number: $BUILD_NUMBER  

Check console output here: $BUILD_URL
"""
        )
    }

    failure {
        emailext(
            to: 'nandhinibalamurugan2003@gmail.com',
            subject: " $JOB_NAME - Build #$BUILD_NUMBER - FAILED",
            body: """
 Build Failed!

Job: $JOB_NAME  
Build Number: $BUILD_NUMBER  

Check what went wrong: $BUILD_URL
"""
        )
    }
}

}
