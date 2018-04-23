package main

import "path/filepath"

type androidAppShared struct {
	androidShared
}

func (aas *androidAppShared) deployableConfig() deployableConfig {
	return nil
}

func (aas *androidAppShared) templateAlias() string {
	return "lang/android_app"
}

func (aas *androidAppShared) GradlePlugins() []string {
	return []string{
		"com.android.application",
	}
}

func (aas *androidAppShared) buildscriptDependencies() []string {
	return []string{
		"com.android.tools.build:gradle",
	}
}

func (aas *androidAppShared) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}

func (aas *androidAppShared) generateExtraRootProjectFiles(data *ProjectData) {
	if data.gdmcRepoURL == "" {
		templateTemplateableToFile("root-gdmc.json", "gdmc.json", aas, data)
	}
}

func (aas *androidAppShared) generateAppResources(data *ProjectData, subdir string) {
	mainRoot := pathWithOptSubdir(subdir, "src", "main")
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
		aas,
		data)
	writeRawTemplateAsset(
		"src_files/res/layout/fragment_main.xml",
		filepath.Join(layout, "fragment_main.xml"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/mipmap-hdpi/ic_launcher.png",
		filepath.Join(mipmapHdpi, "ic_launcher.png"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/mipmap-mdpi/ic_launcher.png",
		filepath.Join(mipmapMdpi, "ic_launcher.png"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/mipmap-xhdpi/ic_launcher.png",
		filepath.Join(mipmapXhdpi, "ic_launcher.png"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/mipmap-xxhdpi/ic_launcher.png",
		filepath.Join(mipmapXxhdpi, "ic_launcher.png"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/mipmap-xxxhdpi/ic_launcher.png",
		filepath.Join(mipmapXxxhdpi, "ic_launcher.png"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/values/colors.xml",
		filepath.Join(values, "colors.xml"),
		aas)
	writeRawTemplateAsset(
		"src_files/res/values/styles.xml",
		filepath.Join(values, "styles.xml"),
		aas)
	templateTemplateableToFile(
		"src_files/res/values/strings.xml",
		filepath.Join(values, "strings.xml"),
		aas,
		data)
}
