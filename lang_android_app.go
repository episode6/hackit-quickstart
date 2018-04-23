package main

import "path/filepath"

type androidApplication struct {
	androidAppShared
}

func (aa *androidApplication) templateAlias() string {
	return "lang/android_app"
}

func (aa *androidApplication) describe() string {
	return "An android application"
}

func (aa *androidApplication) generateLangSpecificFiles(data *ProjectData, subdir string) {
	aa.androidAppShared.generateAppResources(data, subdir)
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "java", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "java", data.Group.asPath())
	androidTestPath := pathWithOptSubdir(subdir, "src", "androidTest", "java", data.Group.asPath())
	mkdir(mainPath, testPath, androidTestPath)

	templateTemplateableToFile(
		"proguard-rules.pro",
		filepath.Join(subdir, "proguard-rules.pro"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/AndroidManifest.xml",
		filepath.Join(mainRoot, "AndroidManifest.xml"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivity.java",
		filepath.Join(mainPath, "MainActivity.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/MainFragment.java",
		filepath.Join(mainPath, "MainFragment.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivityTest.java",
		filepath.Join(testPath, "MainActivityTest.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/MainActivityInstrumentedTest.java",
		filepath.Join(androidTestPath, "MainActivityInstrumentedTest.java"),
		aa,
		data)
}
