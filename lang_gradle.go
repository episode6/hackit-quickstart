package main

import "path/filepath"

type gradleLibrary struct{}

func (gl *gradleLibrary) templateAlias() string {
	return "lang/gradle"
}

func (gl *gradleLibrary) describe() string {
	return "A deployable groovy library with the gradle api and an empty gradle plugin."
}

func (gl *gradleLibrary) deployableConfig() deployableConfig {
	return gl
}

func (gl *gradleLibrary) GradlePlugins() []string {
	return []string{
		"groovy",
		"java-gradle-plugin",
	}
}

func (gl *gradleLibrary) buildscriptRepos() []string {
	return []string{}
}

func (gl *gradleLibrary) projectRepos() []string {
	return []string{}
}

func (gl *gradleLibrary) buildscriptDependencies() []string {
	return []string{}
}

func (gl *gradleLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (gl *gradleLibrary) deployableBuildscriptDependencies() []string {
	return []string{}
}

func (gl *gradleLibrary) deployableGradlePlugins() []string {
	return []string{
		"com.episode6.hackit.deployable.jar",
		"com.episode6.hackit.deployable.addon.groovydocs",
	}
}

func (gl *gradleLibrary) deployableJenkinsCommands() []string {
	return []string{"maybeDeploy"}
}

func (gl *gradleLibrary) generateExtraRootProjectFiles(data *ProjectData) {

}

func (gl *gradleLibrary) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "groovy", data.Group.asPath())
	pluginPath := filepath.Join(mainRoot, "resources", "META-INF", "gradle-plugins")
	testPath := pathWithOptSubdir(subdir, "src", "test", "groovy", data.Group.asPath())
	mkdir(mainPath, pluginPath, testPath)

	templateTemplateableToFile(
		"src_files/SomePlugin.groovy",
		filepath.Join(mainPath, data.CamelName()+"Plugin.groovy"),
		gl,
		data)
	templateTemplateableToFile(
		"src_files/SomePluginTest.groovy",
		filepath.Join(testPath, data.CamelName()+"PluginTest.groovy"),
		gl,
		data)
	templateTemplateableToFile(
		"src_files/plugin.properties",
		filepath.Join(pluginPath, string(data.Group)+".properties"),
		gl,
		data)
}
