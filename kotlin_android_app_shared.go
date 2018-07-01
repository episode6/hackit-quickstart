package main

type kotlinAndroidAppShared struct {
	androidAppShared
}

func (kaa *kotlinAndroidAppShared) templateAlias() string {
	return "lang/kotlin_android_app"
}

func (kaa *kotlinAndroidAppShared) GradlePlugins() []string {
	return append(kaa.androidAppShared.GradlePlugins(),
		"kotlin-android",
		"kotlin-android-extensions")
}

func (kaa *kotlinAndroidAppShared) buildscriptDependencies() []string {
	return append(kaa.androidAppShared.buildscriptDependencies(),
		"org.jetbrains.kotlin:kotlin-gradle-plugin")
}

func (kaa *kotlinAndroidAppShared) generateExtraRootProjectFiles(data *ProjectData) {
	if data.gdmcRepoURL == "" {
		templateTemplateableToFile("root-gdmc.json", "gdmc.json", kaa, data)
	}
}
