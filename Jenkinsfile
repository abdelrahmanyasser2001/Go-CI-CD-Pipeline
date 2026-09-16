pipeline {
    
    agent any
    tools {
        go '1.22.2'
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
                sh 'go build -o main.exe'
            }
        }
        stage('deploy') {
            steps {
                echo 'deploying'
            }
        }
    }
}
