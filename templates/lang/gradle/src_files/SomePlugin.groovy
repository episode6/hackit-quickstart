package {{ .Group }}

import org.gradle.api.Plugin
import org.gradle.api.Project

class {{ .CamelName }}Plugin implements Plugin<Project> {
  void apply(Project project) {
    project.task("helloWorldTask") {
      doLast {
        println "hello world"
      }
    }
  }
}
