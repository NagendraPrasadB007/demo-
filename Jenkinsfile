pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Check out your source code from your version control system (e.g., Git).
                checkout scm
            }
        }

 //stage('Build') {
           // steps {
                // Build your Go application
               // sh "go build -o main"
          //  }
       // }
       

 stage('SonarQube analysis') {
        steps {
            withSonarQubeEnv('SonarQubeServer') { 
                  script {
                      def scannerHome = tool 'sonar';
                      withEnv(["PATH+SONAR=${scannerHome}/bin"]) {
                       sh 'sonar-scanner -Dsonar.projectKey=TRR_Bkend -Dsonar.sources=. -Dsonar.host.url=http://172.28.12.197:9000 -Dsonar.login=sqa_e88af9743e8060253e5188ba77cb699b22c68fcb'
                      }
                   }
         }
       }
    }




       stage('Create Docker Image') {
            steps {
                sh 'docker build -t pickupmanagementbackend:latest .'
            }
        }

        stage('Deploy Docker Image') {
            steps {
                sshagent(['deploy_key']) { //// replace 'my-ssh-key-id' with your actual credential ID
                    sh 'docker save pickupmanagementbackend:latest | ssh admin1@172.28.12.198 "docker load"'
                }
            }
        }

  stage('Stop and Remove Existing Docker Container') {
            steps {
                sshagent(['deploy_key']) {
                    sh 'ssh admin1@172.28.12.198 "docker stop pickupmanagementbackend_container && docker rm pickupmanagementbackend_container || true"'
                }
            }
        }



        stage('Run Docker Container') {
            steps {
                sshagent(['deploy_key']) { // replace 'my-ssh-key-id' with your actual credential ID                 
                 
                    
                    sh '''
                        ssh admin1@172.28.12.198 'docker run -d -p 5450:8080 --name pickupmanagementbackend_container pickupmanagementbackend:latest'

                    '''
                }
            }
        }
    }
}

