package main

import "path/filepath"

type kotlinAndroidLibrary struct {
	androidShared
}

func (kal *kotlinAndroidLibrary) templateAlias() string {
	return "lang/kotlin_android_lib"
}

func (kal *kotlinAndroidLibrary) describe() string {
	return "An android library"
}

func (kal *kotlinAndroidLibrary) deployableConfig() deployableConfig {
	return nil
}

func (kal *kotlinAndroidLibrary) GradlePlugins() []string {
	return []string{
		"com.android.library",
		"kotlin-android",
	}
}

func (kal *kotlinAndroidLibrary) buildscriptDependencies() []string {
	return []string{
		"com.android.tools.build:gradle",
		"org.jetbrains.kotlin:kotlin-gradle-plugin",
	}
}

func (kal *kotlinAndroidLibrary) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (kal *kotlinAndroidLibrary) deployableBuildscriptDependencies() []string {
	return []string{}
}

func (kal *kotlinAndroidLibrary) deployableGradlePlugins() []string {
	return []string{
		"com.episode6.hackit.deployable.aar",
	}
}

func (kal *kotlinAndroidLibrary) deployableJenkinsCommands() []string {
	return []string{"maybeDeploy"}
}

func (kal *kotlinAndroidLibrary) generateExtraRootProjectFiles(data *ProjectData) {
	if data.gdmcRepoURL == "" {
		templateTemplateableToFile("root-gdmc.json", "gdmc.json", kal, data)
	}
}

func (kal *kotlinAndroidLibrary) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "kotlin", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "kotlin", data.Group.asPath())
	mkdir(mainPath, testPath)

	templateTemplateableToFile(
		"proguard-rules.pro",
		filepath.Join(subdir, "proguard-rules.pro"),
		kal,
		data)
	templateTemplateableToFile(
		"src_files/AndroidManifest.xml",
		filepath.Join(mainRoot, "AndroidManifest.xml"),
		kal,
		data)
	templateTemplateableToFile(
		"src_files/SomeClass.kt",
		filepath.Join(mainPath, "SomeClass.kt"),
		kal,
		data)
	templateTemplateableToFile(
		"src_files/SomeClassTest.kt",
		filepath.Join(testPath, "SomeClassTest.kt"),
		kal,
		data)
}
