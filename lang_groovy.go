package main

import "path/filepath"

type groovyLibrary struct{}

func (gl *groovyLibrary) templateAlias() string {
	return "lang/groovy"
}

func (gl *groovyLibrary) describe() string {
	return "A deployable groovy library"
}

func (gl *groovyLibrary) GradlePlugins() []string {
	return []string{
		"groovy",
		"com.episode6.hackit.deployable.jar",
		"com.episode6.hackit.deployable.addon.groovydocs",
		"com.episode6.hackit.gdmc",
	}
}

func (gl *groovyLibrary) buildscriptDependencies() []string {
	return []string{
		"com.episode6.hackit.deployable:deployable",
		"com.episode6.hackit.gdmc:gdmc",
	}
}

func (gl *groovyLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest", "maybeDeploy"}
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
