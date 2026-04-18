pipeline {
    agent any

    stages {
        stage('拉去代码') {
            steps {
                 git branch: 'main', credentialsId: 'ce31a3db-75be-4756-aa68-917e3305a5e0', url: 'git@github.com:lowbroth/demo.git'
            }
        }


           stage('推送代码到远程服务器') {
                    steps {
                      sshPublisher(publishers: [sshPublisherDesc(configName: '192.168.42.100', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '/study/subject', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '**/*')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
                    }
                }
    }
}
