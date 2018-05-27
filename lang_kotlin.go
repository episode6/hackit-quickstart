package main

import (
	"path/filepath"
)

type kotlinLibrary struct{}

func (kl *kotlinLibrary) templateAlias() string {
	return "lang/kotlin"
}

func (kl *kotlinLibrary) describe() string {
	return "A deployable kotlin library"
}

func (kl *kotlinLibrary) deployableConfig() deployableConfig {
	return kl
}

func (kl *kotlinLibrary) GradlePlugins() []string {
	return []string{
		"kotlin",
	}
}

func (kl *kotlinLibrary) buildscriptRepos() []string {
	return []string{}
}

func (kl *kotlinLibrary) projectRepos() []string {
	return []string{}
}

func (kl *kotlinLibrary) buildscriptDependencies() []string {
	return []string{"org.jetbrains.kotlin:kotlin-gradle-plugin"}
}

func (kl *kotlinLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (kl *kotlinLibrary) deployableBuildscriptDependencies() []string {
	return []string{"org.jetbrains.dokka:dokka-gradle-plugin"}
}

func (kl *kotlinLibrary) deployableGradlePlugins() []string {
	return []string{
		"com.episode6.hackit.deployable.kt.jar",
	}
}

func (kl *kotlinLibrary) deployableJenkinsCommands() []string {
	return []string{"maybeDeploy"}
}

func (kl *kotlinLibrary) generateExtraRootProjectFiles(data *ProjectData) {
	if data.gdmcRepoURL == "" {
		templateTemplateableToFile("root-gdmc.json", "gdmc.json", kl, data)
	}
}

func (kl *kotlinLibrary) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainPath := pathWithOptSubdir(subdir, "src", "main", "kotlin", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "kotlin", data.Group.asPath())
	mkdir(mainPath, testPath)

	templateTemplateableToFile(
		"src_files/SomeClass.kt",
		filepath.Join(mainPath, "SomeClass.kt"),
		kl,
		data)
	templateTemplateableToFile(
		"src_files/SomeClassTest.kt",
		filepath.Join(testPath, "SomeClassTest.kt"),
		kl,
		data)
}
