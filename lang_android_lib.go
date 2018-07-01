package main

import "path/filepath"

type androidLibrary struct {
	androidShared
}

func (al *androidLibrary) templateAlias() string {
	return "lang/android_lib"
}

func (al *androidLibrary) describe() string {
	return "An android library"
}

func (al *androidLibrary) deployableConfig() deployableConfig {
	return al
}

func (al *androidLibrary) GradlePlugins() []string {
	return []string{
		"com.android.library",
	}
}

func (al *androidLibrary) buildscriptDependencies() []string {
	return []string{
		"com.android.tools.build:gradle",
	}
}

func (al *androidLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (al *androidLibrary) deployableBuildscriptDependencies() []string {
	return []string{}
}

func (al *androidLibrary) deployableGradlePlugins() []string {
	return []string{
		"com.episode6.hackit.deployable.aar",
	}
}

func (al *androidLibrary) deployableJenkinsCommands() []string {
	return []string{"maybeDeploy"}
}

func (al *androidLibrary) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "java", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "java", data.Group.asPath())
	mkdir(mainPath, testPath)

	templateTemplateableToFile(
		"proguard-rules.pro",
		filepath.Join(subdir, "proguard-rules.pro"),
		al,
		data)
	templateTemplateableToFile(
		"src_files/AndroidManifest.xml",
		filepath.Join(mainRoot, "AndroidManifest.xml"),
		al,
		data)
	templateTemplateableToFile(
		"src_files/SomeClass.java",
		filepath.Join(mainPath, "SomeClass.java"),
		al,
		data)
	templateTemplateableToFile(
		"src_files/SomeClassTest.java",
		filepath.Join(testPath, "SomeClassTest.java"),
		al,
		data)
}
