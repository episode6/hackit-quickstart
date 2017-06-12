package main

type androidShared struct{}

func (as *androidShared) validate(data *ProjectData) {
	require(&data.AndroidCompileSdkVersion, "androidCompileSdkVersion")
	require(&data.AndroidBuildToolsVersion, "androidBuildToolsVersion")
}
