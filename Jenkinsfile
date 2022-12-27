node {
    try {
        currentBuild.result = "SUCCESS"
        env.AWS_ECR_LOGIN = true
        checkout scm
        
        docker.image('golang:1.18').inside('-u root') {
            stage('Setup') {
                echo 'Setting up environment'
                echo env.BRANCH_NAME
                
                if (env.BRANCH_NAME == 'staging') {
                    echo './staging.properties going to load.'
                    configFileProvider([configFile(fileId: 'staging-env-file', targetLocation: './')]) {
                        load './staging.properties'                        
                    }
                    echo 'load properties done.'
                }
                
                if (env.BRANCH_NAME == 'qa') {
                    echo './qa.properties going to load.'
                    configFileProvider([configFile(fileId: 'qa-env-file', targetLocation: './')]) {
                        load './qa.properties'
                    }
                    echo 'load properties done.'
                }
                
                if (env.BRANCH_NAME == 'production') {
                    echo './production.properties going to load.'
                    configFileProvider([configFile(fileId: 'production-env-file', targetLocation: './')]) {
                        load './production.properties'
                    }
                    echo 'load properties done.'
                }
            }
        }
        
        stage('Deploy to qa') {
            echo env.BRANCH_NAME
            if (env.BRANCH_NAME == 'qa') {
                echo 'Building and pushing image'
                docker.withRegistry('https://453230908534.dkr.ecr.ap-south-1.amazonaws.com/tracified/nft-backend', 'ecr:ap-south-1:aws-ecr-credentials') {
                    echo 'Building image'
                    echo "${env.BUILD_ID}"
                    def releaseImage = docker.build("tracified/nft-backend:${env.BUILD_ID}")
                    releaseImage.push()
                    releaseImage.push('latest')
                }
                echo 'Deploying image in server'
                withCredentials([[
                    $class: 'AmazonWebServicesCredentialsBinding',
                    accessKeyVariable: 'AWS_ACCESS_KEY_ID',
                    credentialsId: 'aws-ecr-credentials',
                    secretKeyVariable: 'AWS_SECRET_ACCESS_KEY'
                ]]) {
                    ansiblePlaybook inventory: 'deploy/hosts', playbook: 'deploy/qa.yml', extras: '-u ubuntu -e GATEWAY_PORT=$GATEWAY_PORT'
                }
            }
        }
        
        stage('Deploy to Production') {
            echo env.BRANCH_NAME
            if (env.BRANCH_NAME == 'master') {
                echo 'Building and pushing image'
                docker.withRegistry('https://453230908534.dkr.ecr.ap-south-1.amazonaws.com/tracified/nft-backend-production', 'ecr:ap-south-1:aws-ecr-credentials') {
                    echo 'Building image'
                    echo "${env.BUILD_ID}"                  
                    def releaseImage = docker.build("tracified/nft-backend-production:${env.BUILD_ID}")
                    releaseImage.push()
                    releaseImage.push('latest')
                }
            }
        }
    }
    catch (exc) {
        currentBuild.result = "FAILURE"
        echo 'Something went wrong'
        echo exc.toString()
    }
    finally {
        echo 'All done. Cleaning up docker'
        sh 'docker system prune -af'
        echo "Completed pipeline: ${currentBuild.fullDisplayName} with status of ${currentBuild.result}"
    }
}
