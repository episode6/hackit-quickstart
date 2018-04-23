package main

// rootProject is a partial implementation of projectTemplate
type rootProject struct{}

func (rp *rootProject) validate(data *ProjectData) {
	require(&data.Version, "version")
}

func (rp *rootProject) generate(data *ProjectData) {
	if data.gdmcRepoURL != "" {
		addGitSubmodule(data.gdmcRepoURL, "gdmc")
	}
	templateAssetToFile("gitignore", ".gitignore", data)
	templateAssetToFile("root-build.gradle", "build.gradle", data)
	templateTemplateableToFile("settings.gradle", "settings.gradle", data.Proj, data)
	templateTemplateableToFile("proj-gradle.properties", "gradle.properties", data.Lang, data)
	templateTemplateableToFile("proj-local.properties", "local.properties", data.Lang, data)
	templateAssetToFile("Jenkinsfile", "Jenkinsfile", data)
	data.Lang.generateExtraRootProjectFiles(data)
	execOrPanic("gradle -Dorg.gradle.daemon=false -Pgdmc.forceResolve=true wrapper")
}
