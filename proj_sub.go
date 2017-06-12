package main

import (
	"path/filepath"
)

type subProject struct {
	skipAmmendSettings bool
}

func (sp *subProject) templateAlias() string {
	return "proj/sub"
}

func (sp *subProject) describe() string {
	return "A new submodule in an existing multi-module project"
}

func (sp *subProject) generate(data *ProjectData) {
	mkdir(data.Name)
	templateAssetToFile(
		"proj-build.gradle",
		filepath.Join(data.Name, "build.gradle"),
		data)
	if !sp.skipAmmendSettings {
		appendTextFile("settings.gradle", ", ':"+data.Name+"'")
	}
	data.Lang.generateLangSpecificFiles(data, data.Name)
	execGdmcResolve()
}
