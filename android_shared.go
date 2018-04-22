package main

type androidShared struct{}

func (as *androidShared) validate(data *ProjectData) {
	require(&data.AndroidSdkDir, "androidSdkDir")
	require(&data.AndroidCompileSdkVersion, "androidCompileSdkVersion")
}

func (as *androidShared) generateExtraRootProjectFiles(data *ProjectData) {

}

func (as *androidShared) buildscriptRepos() []string {
	return []string{"google()"}
}

func (as *androidShared) projectRepos() []string {
	return []string{"google()"}
}
