pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go install'
            }
        }
        stage('test') {
            steps {
                sh 'go test -v'
            }
        }
    }
}
