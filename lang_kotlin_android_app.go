package main

import "path/filepath"

type kotlinAndroidApplication struct {
	androidAppShared
}

func (kaa *kotlinAndroidApplication) templateAlias() string {
	return "lang/kotlin_android_app"
}

func (kaa *kotlinAndroidApplication) describe() string {
	return "An android application with kotlin support"
}

func (kaa *kotlinAndroidApplication) GradlePlugins() []string {
	return append(kaa.androidAppShared.GradlePlugins(),
		"kotlin-android",
		"kotlin-android-extensions")
}

func (kaa *kotlinAndroidApplication) buildscriptDependencies() []string {
	return append(kaa.androidAppShared.buildscriptDependencies(),
		"org.jetbrains.kotlin:kotlin-gradle-plugin")
}

func (kaa *kotlinAndroidApplication) generateExtraRootProjectFiles(data *ProjectData) {
	if data.gdmcRepoURL == "" {
		templateTemplateableToFile("root-gdmc.json", "gdmc.json", kaa, data)
	}
}

func (kaa *kotlinAndroidApplication) generateLangSpecificFiles(data *ProjectData, subdir string) {
	kaa.androidAppShared.generateAppResources(data, subdir)
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "java", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "java", data.Group.asPath())
	androidTestPath := pathWithOptSubdir(subdir, "src", "androidTest", "java", data.Group.asPath())
	mkdir(mainPath, testPath, androidTestPath)

	templateTemplateableToFile(
		"proguard-rules.pro",
		filepath.Join(subdir, "proguard-rules.pro"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/AndroidManifest.xml",
		filepath.Join(mainRoot, "AndroidManifest.xml"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivity.java",
		filepath.Join(mainPath, "MainActivity.java"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainFragment.java",
		filepath.Join(mainPath, "MainFragment.java"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivityTest.java",
		filepath.Join(testPath, "MainActivityTest.java"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivityInstrumentedTest.java",
		filepath.Join(androidTestPath, "MainActivityInstrumentedTest.java"),
		kaa,
		data)
}
