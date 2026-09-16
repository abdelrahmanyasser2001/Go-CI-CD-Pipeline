pipeline {
    
    agent any
    tools {
        go '1.22.2'
    }
    environment {
    DOCKER_CREDS = credentials('docker-hub-credentials')
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
                echo 'building the artifact and pushing to docker hub'
                sh 'echo $DOCKER_HUB_PSW| docker login -u $DOCKER_CREDS_USR --password-stdin'
                sh 'docker build -t abdelrahmanyasserhub/go-app:1.0 .'
                sh 'docker push abdelrahmanyasserhub/go-app:1.0'
           }
        }
        stage('deploy') {
            steps {
                echo 'deploying'
            }
        }
    }
}


