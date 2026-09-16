pipeline {
    agent any

    stages {
        stage('test') {
            steps {
                echo 'testing the application'
                sh 'go test -v ./test'
            }
        }
        stage('build') {
            steps {
                echo 'building the application'
            }
        }
        stage('deploy') {
            steps {
                echo 'deploying'
            }
        }
    }
}
