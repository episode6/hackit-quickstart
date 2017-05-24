package main

import "path/filepath"

type androidApplication struct{}

func (aa *androidApplication) templateAlias() string {
	return "lang/android_app"
}

func (aa *androidApplication) describe() string {
	return "An android application"
}

func (aa *androidApplication) GradlePlugins() []string {
	return []string{
		"com.android.application",
		"com.episode6.hackit.gdmc",
	}
}

func (aa *androidApplication) buildscriptDependencies() []string {
	return []string{
		"com.android.tools.build:gradle",
		"com.episode6.hackit.gdmc:gdmc",
	}
}

func (aa *androidApplication) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (aa *androidApplication) generateLangSpecificFiles(data *ProjectData, subdir string) {
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

	writeRawTemplateAsset(
		"src_files/res/layout/activity_main.xml",
		filepath.Join(layout, "activity_main.xml"),
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
