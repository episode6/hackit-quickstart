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

func (kad *kotlinAndroidApplicationWithDagger) generateLangSpecificFiles(data *ProjectData, subdir string) {
	kad.androidAppShared.generateAppResources(data, subdir)
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "java", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "java", data.Group.asPath())
	androidTestPath := pathWithOptSubdir(subdir, "src", "androidTest", "java", data.Group.asPath())
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
		"src_files/app/App.java",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"App.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/app/AppComponent.java",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"AppComponent.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/app/AppModule.java",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"AppModule.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/app/RootBindingModule.java",
		filepath.Join(mainAppPath, "RootBindingModule.java"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/base/BaseAppCompatActivity.java",
		filepath.Join(mainBasePath, "BaseAppCompatActivity.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/base/BaseFragment.java",
		filepath.Join(mainBasePath, "BaseFragment.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/base/BaseActivityModule.java",
		filepath.Join(mainBasePath, "BaseActivityModule.java"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/main/MainActivity.java",
		filepath.Join(mainMainPath, "MainActivity.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/main/MainFragment.java",
		filepath.Join(mainMainPath, "MainFragment.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/main/MainActivityModule.java",
		filepath.Join(mainMainPath, "MainActivityModule.java"),
		kad,
		data)

	templateTemplateableToFile(
		"src_files/test/ut/BuildMockspresso.java",
		filepath.Join(testPath, "BuildMockspresso.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/main/MainActivityTest.java",
		filepath.Join(testMainPath, "MainActivityTest.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/main/MainFragmentTest.java",
		filepath.Join(testMainPath, "MainFragmentTest.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/app/MockspressoTestApp.java",
		filepath.Join(testAppPath, "MockspressoTestApp.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/base/BaseAppCompatActivityTest.java",
		filepath.Join(testBasePath, "BaseAppCompatActivityTest.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/ut/base/BaseFragmentTest.java",
		filepath.Join(testBasePath, "BaseFragmentTest.java"),
		kad,
		data)
	templateTemplateableToFile(
		"src_files/test/it/MainActivityInstrumentedTest.java",
		filepath.Join(androidTestPath, "MainActivityInstrumentedTest.java"),
		kad,
		data)
}
