package main

import "path/filepath"

type kotlinAndroidApplication struct {
	kotlinAndroidAppShared
}

func (kaa *kotlinAndroidApplication) templateAlias() string {
	return "lang/kotlin_android_app"
}

func (kaa *kotlinAndroidApplication) describe() string {
	return "An android application with kotlin support"
}

func (kaa *kotlinAndroidApplication) generateLangSpecificFiles(data *ProjectData, subdir string) {
	kaa.androidAppShared.generateAppResources(data, subdir)
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "kotlin", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "kotlin", data.Group.asPath())
	androidTestPath := pathWithOptSubdir(subdir, "src", "androidTest", "kotlin", data.Group.asPath())
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
		"src_files/MainActivity.kt",
		filepath.Join(mainPath, "MainActivity.kt"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainFragment.kt",
		filepath.Join(mainPath, "MainFragment.kt"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivityTest.kt",
		filepath.Join(testPath, "MainActivityTest.kt"),
		kaa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivityInstrumentedTest.kt",
		filepath.Join(androidTestPath, "MainActivityInstrumentedTest.kt"),
		kaa,
		data)
}
