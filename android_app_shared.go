package main

type androidAppShared struct {
	androidShared
}

func (aas *androidAppShared) GradlePlugins() []string {
	return []string{
		"com.android.application",
		"com.episode6.hackit.gdmc",
	}
}

func (aas *androidAppShared) buildscriptDependencies() []string {
	return []string{
		"com.android.tools.build:gradle",
		"com.episode6.hackit.gdmc:gdmc",
	}
}

func (aas *androidAppShared) JenkinsCommands() []string {
	return []string{"buildAndTest"}
}
