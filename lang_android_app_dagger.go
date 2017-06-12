package main

import "path/filepath"

type androidApplicationWithDagger struct {
	androidAppShared
}

func (aa *androidApplicationWithDagger) templateAlias() string {
	return "lang/android_app_dagger"
}

func (aa *androidApplicationWithDagger) describe() string {
	return "An android application with a default dagger 2 implementation"
}

func (aa *androidApplicationWithDagger) generateLangSpecificFiles(data *ProjectData, subdir string) {
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
	mainPath := filepath.Join(mainRoot, "java", data.Group.asPath())
	testPath := pathWithOptSubdir(subdir, "src", "test", "java", data.Group.asPath())
	androidTestPath := pathWithOptSubdir(subdir, "src", "androidTest", "java", data.Group.asPath())
	mainAppPath := filepath.Join(mainPath, "app")
	mainMainPath := filepath.Join(mainPath, "main")
	mainBasePath := filepath.Join(mainPath, "base")
	mkdir(mainPath, testPath, androidTestPath, mainAppPath, mainMainPath, mainBasePath)

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
		"src_files/app/App.java",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"App.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/app/AppComponent.java",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"AppComponent.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/app/AppModule.java",
		filepath.Join(mainAppPath, data.CamelNameWithoutApp()+"AppModule.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/app/RootBindingModule.java",
		filepath.Join(mainAppPath, "RootBindingModule.java"),
		aa,
		data)

	templateTemplateableToFile(
		"src_files/base/BaseAppCompatActivity.java",
		filepath.Join(mainBasePath, "BaseAppCompatActivity.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/base/BaseFragment.java",
		filepath.Join(mainBasePath, "BaseFragment.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/base/BaseActivityModule.java",
		filepath.Join(mainBasePath, "BaseActivityModule.java"),
		aa,
		data)

	templateTemplateableToFile(
		"src_files/main/MainActivity.java",
		filepath.Join(mainMainPath, "MainActivity.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/main/MainFragment.java",
		filepath.Join(mainMainPath, "MainFragment.java"),
		aa,
		data)
	templateTemplateableToFile(
		"src_files/main/MainActivityModule.java",
		filepath.Join(mainMainPath, "MainActivityModule.java"),
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

	resPath := filepath.Join(mainRoot, "res")
	layout := filepath.Join(resPath, "layout")
	mipmapHdpi := filepath.Join(resPath, "mipmap-hdpi")
	mipmapMdpi := filepath.Join(resPath, "mipmap-mdpi")
	mipmapXhdpi := filepath.Join(resPath, "mipmap-xhdpi")
	mipmapXxhdpi := filepath.Join(resPath, "mipmap-xxhdpi")
	mipmapXxxhdpi := filepath.Join(resPath, "mipmap-xxxhdpi")
	values := filepath.Join(resPath, "values")
	mkdir(layout, mipmapHdpi, mipmapMdpi, mipmapXhdpi, mipmapXxhdpi, mipmapXxxhdpi, values)

	templateTemplateableToFile(
		"src_files/res/layout/activity_main.xml",
		filepath.Join(layout, "activity_main.xml"),
		aa,
		data)
	writeRawTemplateAsset(
		"src_files/res/layout/fragment_main.xml",
		filepath.Join(layout, "fragment_main.xml"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/mipmap-hdpi/ic_launcher.png",
		filepath.Join(mipmapHdpi, "ic_launcher.png"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/mipmap-mdpi/ic_launcher.png",
		filepath.Join(mipmapMdpi, "ic_launcher.png"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/mipmap-xhdpi/ic_launcher.png",
		filepath.Join(mipmapXhdpi, "ic_launcher.png"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/mipmap-xxhdpi/ic_launcher.png",
		filepath.Join(mipmapXxhdpi, "ic_launcher.png"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/mipmap-xxxhdpi/ic_launcher.png",
		filepath.Join(mipmapXxxhdpi, "ic_launcher.png"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/values/colors.xml",
		filepath.Join(values, "colors.xml"),
		aa)
	writeRawTemplateAsset(
		"src_files/res/values/styles.xml",
		filepath.Join(values, "styles.xml"),
		aa)
	templateTemplateableToFile(
		"src_files/res/values/strings.xml",
		filepath.Join(values, "strings.xml"),
		aa,
		data)
}
