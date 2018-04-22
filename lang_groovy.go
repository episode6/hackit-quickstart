package main

import "path/filepath"

type groovyLibrary struct{}

func (gl *groovyLibrary) templateAlias() string {
	return "lang/groovy"
}

func (gl *groovyLibrary) describe() string {
	return "A deployable groovy library"
}

func (gl *groovyLibrary) deployableConfig() deployableConfig {
	return gl
}

func (gl *groovyLibrary) GradlePlugins() []string {
	return []string{
		"groovy",
	}
}

func (gl *groovyLibrary) buildscriptRepos() []string {
	return []string{}
}

func (gl *groovyLibrary) projectRepos() []string {
	return []string{}
}

func (gl *groovyLibrary) buildscriptDependencies() []string {
	return []string{}
}

func (gl *groovyLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (gl *groovyLibrary) deployableBuildscriptDependencies() []string {
	return []string{}
}

func (gl *groovyLibrary) deployableGradlePlugins() []string {
	return []string{
		"com.episode6.hackit.deployable.jar",
		"com.episode6.hackit.deployable.addon.groovydocs",
	}
}

func (gl *groovyLibrary) deployableJenkinsCommands() []string {
	return []string{"maybeDeploy"}
}

func (gl *groovyLibrary) generateExtraRootProjectFiles(data *ProjectData) {

}

func (gl *groovyLibrary) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainPath := pathWithOptSubdir(subdir, "src", "main", "groovy", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "groovy", data.Group.asPath())
	mkdir(mainPath, testPath)

	templateTemplateableToFile(
		"src_files/SomeClass.groovy",
		filepath.Join(mainPath, "SomeClass.groovy"),
		gl,
		data)
	templateTemplateableToFile(
		"src_files/SomeClassTest.groovy",
		filepath.Join(testPath, "SomeClassTest.groovy"),
		gl,
		data)
}
