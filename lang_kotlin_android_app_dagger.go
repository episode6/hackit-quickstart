package main

import "path/filepath"

type kotlinAndroidApplicationWithDagger struct {
	kotlinAndroidAppShared
}

func (kad *kotlinAndroidApplicationWithDagger) templateAlias() string {
	return "lang/kotlin_android_app_dagger"
}

func (kad *kotlinAndroidApplicationWithDagger) describe() string {
	return "An android application with kotlin support and a default dagger 2 implementation"
}

func (kad *kotlinAndroidApplicationWithDagger) GradlePlugins() []string {
	return append(kad.kotlinAndroidAppShared.GradlePlugins(), "kotlin-kapt")
}

func (kad *kotlinAndroidApplicationWithDagger) generateLangSpecificFiles(data *ProjectData, subdir string) {
	kad.androidAppShared.generateAppResources(data, subdir)
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "kotlin", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "kotlin", data.Group.asPath())
	androidTestPath := pathWithOptSubdir(subdir, "src", "androidTest", "kotlin", data.Group.asPath())
	mainAppPath := filepath.Join(mainPath, "app")
	mainMainPath := filepath.Join(mainPath, "main")
	mainBasePath := filepath.Join(mainPath, "base")
	testAppPath := filepath.Join(testPath, "app")
	testMainPath := filepath.Join(testPath, "main")
	testBasePath := filepath.Join(testPath, "base")
	mkdir(androidTestPath, mainAppPath, mainMainPath, mainBasePath, testAppPath, testMainPath, testBasePath)

	templateTemplateableToFile(
		"proguard-rules.pro",
		filepath.Join(subdir, "proguard-rules.pro"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/AndroidManifest.xml",
		filepath.Join(mainRoot, "AndroidManifest.xml"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/app/App.kt",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"App.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/app/AppComponent.kt",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"AppComponent.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/app/AppModule.kt",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"AppModule.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/app/RootBindingModule.kt",
		filepath.Join(mainAppPath, "RootBindingModule.kt"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/base/BaseAppCompatActivity.kt",
		filepath.Join(mainBasePath, "BaseAppCompatActivity.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/base/BaseFragment.kt",
		filepath.Join(mainBasePath, "BaseFragment.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/base/BaseActivityModule.kt",
		filepath.Join(mainBasePath, "BaseActivityModule.kt"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/main/MainActivity.kt",
		filepath.Join(mainMainPath, "MainActivity.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/main/MainFragment.kt",
		filepath.Join(mainMainPath, "MainFragment.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/main/MainActivityModule.kt",
		filepath.Join(mainMainPath, "MainActivityModule.kt"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/test/ut/BuildMockspresso.kt",
		filepath.Join(testPath, "BuildMockspresso.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/main/MainActivityTest.kt",
		filepath.Join(testMainPath, "MainActivityTest.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/main/MainFragmentTest.kt",
		filepath.Join(testMainPath, "MainFragmentTest.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/app/MockspressoTestApp.kt",
		filepath.Join(testAppPath, "MockspressoTestApp.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/base/BaseAppCompatActivityTest.kt",
		filepath.Join(testBasePath, "BaseAppCompatActivityTest.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/base/BaseFragmentTest.kt",
		filepath.Join(testBasePath, "BaseFragmentTest.kt"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/it/MainActivityInstrumentedTest.kt",
		filepath.Join(androidTestPath, "MainActivityInstrumentedTest.kt"),
		kad,
		data)
}
