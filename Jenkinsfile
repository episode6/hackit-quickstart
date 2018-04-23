node {
  def srcDir = 'src/github.com/episode6/hackit-quickstart'
  def goName = 'Go 1.8.1'
  stage('checkout') {
    dir(srcDir) {
      checkout scm
      sh 'git submodule update --init'
    }
  }

  def goRunner
  stage('pipeline') {
    goRunner = fileLoader.fromGit(
        'go/GoRunner',
        'git@github.com:episode6/jenkins-pipelines.git',
        'v0.0.7',
        null,
        '')
  }

  def gradleRoot = tool name: 'Gradle 4.4', type: 'gradle'
  withEnv(["PATH+=:${gradleRoot}/bin"]) {
    goRunner.buildAndTest(srcDir, goName, "480m")
  }
}
