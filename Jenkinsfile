pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go install'
            }
        }
        stage('build') {
            steps {
                sh 'go test -v'
            }
        }
    }
}
