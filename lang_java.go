package main

import (
	"path/filepath"
)

type javaLibrary struct{}

func (jl *javaLibrary) templateAlias() string {
	return "lang/java"
}

func (jl *javaLibrary) describe() string {
	return "A deployable java library"
}

func (jl *javaLibrary) deployableConfig() deployableConfig {
	return nil
}

func (jl *javaLibrary) GradlePlugins() []string {
	return []string{
		"java-library",
		"com.episode6.hackit.deployable.jar",
		"com.episode6.hackit.gdmc",
	}
}

func (jl *javaLibrary) buildscriptDependencies() []string {
	return []string{
		"com.episode6.hackit.deployable:deployable",
		"com.episode6.hackit.gdmc:gdmc",
	}
}

func (jl *javaLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest", "maybeDeploy"}
}

func (jl *javaLibrary) generateExtraRootProjectFiles(data *ProjectData) {

}

func (jl *javaLibrary) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainPath := pathWithOptSubdir(subdir, "src", "main", "java", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "java", data.Group.asPath())
	mkdir(mainPath, testPath)

	templateTemplateableToFile(
		"src_files/SomeClass.java",
		filepath.Join(mainPath, "SomeClass.java"),
		jl,
		data)
	templateTemplateableToFile(
		"src_files/SomeClassTest.java",
		filepath.Join(testPath, "SomeClassTest.java"),
		jl,
		data)
}
