pipeline {
    
    agent any
    tools {
        go {
            version '1.22.2'
        }
    }
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
