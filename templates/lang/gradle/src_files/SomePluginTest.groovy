package {{ .Group }}

import org.gradle.testkit.runner.GradleRunner
import org.gradle.testkit.runner.TaskOutcome
import org.junit.Rule
import org.junit.rules.TemporaryFolder
import spock.lang.Specification

class {{ .CamelName }}PluginTest extends Specification {

  @Rule final TemporaryFolder rootDir = new TemporaryFolder()

  def "test init"() {
    given:
    rootDir.newFile("build.gradle") << """
plugins {
  id '{{ .Group }}'
}
"""

    when:
    def result = GradleRunner.create()
        .withProjectDir(rootDir.root)
        .withPluginClasspath()
        .withArguments("helloWorldTask")
        .build()

    then:
    result.task(":helloWorldTask").outcome == TaskOutcome.SUCCESS
    result.output.contains("hello world")
  }
}
