pipeline{
    agent any
    environment {
            PROJECT_PATH = "/go/src/github.com/cjburchell/profilux-go"
    }

        stages {
            stage('Setup') {
                steps {
                    script{
                        slackSend color: "good", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} started"
                    }
                 /* Let's make sure we have the repository cloned to our workspace */
                 checkout scm
                 }
             }
            stage('Static Analysis') {
                parallel {
                    stage('Vet') {
                        agent {
                            docker {
                                image 'cjburchell/goci:1.14'
                                args '-v $WORKSPACE:$PROJECT_PATH'
                            }
                        }
                        steps {
                            script{
                                    sh """go vet ./..."""
                                    def checkVet = scanForIssues tool: [$class: 'GoVet']
                                    publishIssues issues:[checkVet]
                            }
                        }
                    }

                    stage('Lint') {
                        agent {
                            docker {
                                image 'cjburchell/goci:1.14'
                                args '-v $WORKSPACE:$PROJECT_PATH'
                            }
                        }
                        steps {
                            script{
                                sh """golint ./..."""

                                def checkLint = scanForIssues tool: [$class: 'GoLint']
                                publishIssues issues:[checkLint]
                            }
                        }
                    }
                }
            }
        }

        post {
            always {
                  script{
                      if ( currentBuild.currentResult == "SUCCESS" ) {
                        slackSend color: "good", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was successful"
                      }
                      else if( currentBuild.currentResult == "FAILURE" ) {
                        slackSend color: "danger", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was failed"
                      }
                      else if( currentBuild.currentResult == "UNSTABLE" ) {
                        slackSend color: "warning", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was unstable"
                      }
                      else {
                        slackSend color: "danger", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} its result (${currentBuild.currentResult}) was unclear"
                      }
                  }
            }
        }
    }